package sandbox_runner

import (
	"executor/pkg/compiler60_shared"
	"golang.org/x/sys/unix"
	"time"
)

const (
	// UnshareFlags is flags used to create namespaces
	// NS - for mounts
	// not USER - cause for file access we want to use the random user id we setuid
	// not CGROUP - cause cgroup fs is not mounted
	// not PID - cause syscalls do not allow any pid access (e.g. kill, ptrace)
	// not NET - cause syscalls do not allow any network access
	// not IPC - cause syscalls do not allow any ipc access
	// not UTS - cause syscalls do not allow any uname, domainname access
	UnshareFlags = unix.CLONE_NEWNS

	MaxConcurrentRunners = 32

	// kill after 4s using context.WithTimeout
	MaxRuntime = 4 * time.Second

	// for cpu time rlimit cannot go below 1 sec
	// we actually want 0.2s cpu time for that we use cgroups with cpu quota and kill after 4s (MaxRuntime)
	RlimitsCpuMax = 1

	RlimitsData      = 128 * 1024 // should mostly be bss
	RlimitsStack     = 300 * 1024
	RlimitsAddrSpace = compiler60_shared.MaxBinarySize + RlimitsData + RlimitsStack // add up the other limits

	RlimitsFileSize  = 128 // this limits how big files can be created
	RLimitsFileCount = 5   // at most 10 fds open - cause we do not provide the close() syscall this also limits the file count overall

	// limit cpu usage 5% each sec
	// cause killed after 4s (MaxRuntime) this should also limit cpu usage to 0.2s per run
	CGroupCPUQuota  = uint64(50_000)    // in us
	CGroupCPUPeriod = uint64(1_000_000) // in us
	CGroupMaxMem    = RlimitsAddrSpace

	CloseFd = ^uintptr(0) // uintptr(-1) for closing fd in forkexec runner
)
