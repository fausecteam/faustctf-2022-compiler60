package sandbox_runner

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"executor/pkg/compiler60_shared"
	"fmt"
	"github.com/criyle/go-sandbox/pkg/cgroup"
	"github.com/criyle/go-sandbox/pkg/forkexec"
	"github.com/criyle/go-sandbox/pkg/mount"
	"github.com/criyle/go-sandbox/pkg/rlimit"
	"golang.org/x/sys/unix"
	"math/rand"
	"os"
	"sync"
	"syscall"
)

var rLimits = (&rlimit.RLimits{
	CPU:          RlimitsCpuMax,
	CPUHard:      RlimitsCpuMax,
	Data:         RlimitsData,
	Stack:        RlimitsStack,
	AddressSpace: RlimitsAddrSpace,

	FileSize: RlimitsFileSize,
	OpenFile: RLimitsFileCount,

	DisableCore: true,
}).PrepareRLimit()

var cgroupNestingOnce sync.Once

func configureNewCgroup() (cg cgroup.Cgroup, err error) {
	cgroupNestingOnce.Do(func() {
		// in docker "cgroup.subtree_control" is not populated, so subgroups would be useless
		if os.Getpid() != 1 {
			return // not in docker (with pid namespace)
		}
		// this is the workaround described here: https://unix.stackexchange.com/questions/683012/how-can-cgroup-subtree-control-be-populated-in-dockers-private-cgroup-namespace
		if err := cgroup.EnableV2Nesting(); err != nil {
			fmt.Printf("Broken cgroup %s\n", err)
		}
	})

	// setup cgroup
	cg, err = cgroup.NewBuilder("compiler60-executor").
		WithType(cgroup.CgroupTypeV2).
		WithCPU().
		WithMemory().
		Random("compiler60-executor")
	if err != nil {
		return
	}
	if err = cg.SetCPUBandwidth(CGroupCPUQuota, CGroupCPUPeriod); err != nil {
		return
	}
	if err = cg.SetMemoryLimit(CGroupMaxMem); err != nil {
		return
	}

	return
}

func (r *Runner) setupRunner() (forkRunner *forkexec.Runner, err error) {
	if len(r.ExecFile) > compiler60_shared.MaxBinarySize {
		err = fmt.Errorf("binary to large %d", len(r.ExecFile))
		return
	}

	mountBuilder := mount.NewBuilder().
		WithBind("/data", "data", false).
		// bind /proc/self/fd for execve("/proc/self/fd/%d")
		WithBind("/proc/self/fd", "proc/self/fd", true)
	if IsDebug() {
		// Write file to 'tmp' directory (/ cause /tmp is not mounted in docker container).
		// And then bind mount it into mount namespace at /binary.
		// Then later you can use "file target:/binary" in gdb.
		// The file is deleted in the SyncFunction via cleanupAtForkSync().
		file, err := os.CreateTemp("/", "compiler60-binary")
		if err == nil {
			r.execFileTmpPath = file.Name()
			_, err = file.Write(r.ExecFile)
			if err == nil {
				mountBuilder.WithBind(file.Name(), "binary", true)
			}
		}
		if err != nil {
			fmt.Printf("Setting up debugging failed: %s\n", err)
		}
	}
	mounts, err := mountBuilder.Build()
	if err != nil {
		return
	}

	// wrap binary bytes in memfd
	err = r.createMemFDOfBinary()
	if err != nil {
		return
	}

	// create pipe for stdout
	err = r.createStdoutPipe()
	if err != nil {
		return
	}

	cred := generateUIDGIDForExec()

	syncFunction := func(pid int) error {
		// close memfd and write side of pipe in parent after fork
		r.cleanupAtForkSync()

		// add forked process to cgroup
		return r.cg.AddProc(pid)
	}

	// ideally we would just use
	// ExecFile:   r.execFileFd.Fd()
	// in the forkexec.Runner so it would use execveat().
	// However the debian bullseye kernel has a bug which makes all fds not work with execveat() after a pivot_root()
	// But execve("/proc/self/fd/%d") still works
	// But, then there is the seccomp filter. We need to allow the first execve, but deny all others.
	// So I patched go-sanbox to pass the execve-path as *byte
	// and set the seccomp filter to only allow this specific pointer value as path.
	// As go places its heap at about 0xc000000000 this is not accessible without mmap.
	execFilePathStr := fmt.Sprintf("/proc/self/fd/%d", r.execFileFd.Fd())
	execFilePath, err := unix.BytePtrFromString(execFilePathStr)
	if err != nil {
		return
	}
	// seccomp policy to allow only read, write and open syscalls
	// allow execve only for pointer value of execFilePath
	seccompPolicy := GetSeccompPolicy(execFilePath)
	forkRunner = &forkexec.Runner{
		ExecPath: execFilePath,

		RLimits:    rLimits,
		Files:      []uintptr{CloseFd, uintptr(r.outPipeWriteFd), CloseFd},
		WorkDir:    "/data",
		Seccomp:    seccompPolicy,
		NoNewPrivs: true,
		CloneFlags: UnshareFlags,
		PivotRoot:  "/etc",
		Mounts:     mounts,
		DropCaps:   true,
		SyncFunc:   syncFunction,
		Credential: &cred,

		// Do a kill(SIGSTOP) in the client to pause execution and allow debugging
		StopBeforeSeccomp: IsDebug(),
		// we do not really want to ptrace but otherwise forkexec_runner.Start() does not return.
		// The logic is a bit tangled in the forkexec_runner:
		// With Ptrace first the child sync is done then kill(SIGSTOP).
		// Without Ptrace first kill(SIGSTOP) is done then sync with parent.
		// We want the first one so forkexec_runner.Start() returns without waiting for the continue.
		// Then we are able to kill the PID if the HTTP request is canceled.
		// Because the child now calls ptrace(PTRACE_TRACEME) we have to detach from it.
		// This also has the benefit that we can continue the execution from the kill(SIGSTOP)
		// until the exec-call, so users cannot/do not have to debug the child.
		Ptrace: IsDebug(),
	}

	return forkRunner, nil
}

var seedRandOnce sync.Once

func generateUIDGIDForExec() syscall.Credential {
	seedRandOnce.Do(func() {
		var b [8]byte
		_, err := cryptorand.Read(b[:])
		if err != nil {
			rand.Seed(42)
		} else {
			rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
		}
	})
	for {
		val := rand.Uint32()
		if val > 1000 && val < (^uint32(0)) { // uint32(-1)
			return syscall.Credential{
				Uid:    val,
				Gid:    val,
				Groups: []uint32{},
			}
		}
	}
}
