package sandbox_runner

import (
	"context"
	"executor/pkg/compiler60_shared"
	"fmt"
	"github.com/criyle/go-sandbox/pkg/forkexec"
	"github.com/criyle/go-sandbox/runner"
	"golang.org/x/sys/unix"
	"io"
	"time"
)

type readResult struct {
	Bytes []byte
	Error error
}

func (r *Runner) startAndWait(ctx context.Context, forkRunner *forkexec.Runner) (result RunnerResult, err error) {
	// Start the runner
	pid, err := forkRunner.Start()
	if err != nil {
		return
	}

	if IsDebug() {
		err = setupChildForExternalDebugging(pid, forkRunner)
		if err != nil {
			fmt.Printf("Setting up debugging failed: %s\n", err)
		}
	}

	// kill all tracee upon return
	defer killAndCollectZombies(pid)

	stdoutReadChan := make(chan readResult)

	go func() {
		var buffer [compiler60_shared.MaxStdoutSize + 1]byte
		n, readErr := io.ReadFull(r.outPipeRead, buffer[:])
		if n == 0 && readErr != io.EOF && readErr != io.ErrUnexpectedEOF {
			stdoutReadChan <- readResult{nil, readErr}
		} else {
			stdoutReadChan <- readResult{buffer[:n], nil}
		}
	}()

	stdoutTooLarge := false
	select {
	case readResult := <-stdoutReadChan:
		if readResult.Error != nil {
			result.Status = runner.StatusRunnerError
			result.Error = readResult.Error
			return
		}
		if len(readResult.Bytes) > compiler60_shared.MaxStdoutSize {
			stdoutTooLarge = true
		} else {
			result.Stdout = readResult.Bytes
		}
	case <-ctx.Done():
		// exceeded MaxRuntime
		if IsDebug() {
			// MaxRuntime is disabled so must be request closed
			fmt.Println("Request canceled - killing debuggee")
		}
	}

	// make sure child will die soon
	killAll(pid)

	var (
		wstatus unix.WaitStatus // wait4 wait status
		rusage  unix.Rusage     // wait4 rusage
	)

	// TODO cloud use signal.Notify(SIGCHLD) to wait for child exit efficiently

	wpid, err := unix.Wait4(pid, &wstatus, 0, &rusage)
	if err != nil {
		return
	}

	// store resource usage
	result.Time = time.Duration(rusage.Utime.Nano()) // ns

	if wpid != pid {
		result.Status = runner.StatusTimeLimitExceeded
	}
	if stdoutTooLarge {
		result.Status = runner.StatusOutputLimitExceeded
	}
	result.WStatus = wstatus
	return
}

// kill all tracee according to pids
func killAll(pid int) {
	_ = unix.Kill(-pid, unix.SIGKILL) // ignore err
}

// collect died child processes
func killAndCollectZombies(pid int) {
	killAll(pid)
	var wstatus unix.WaitStatus
	for {
		// wait for ECHILD from wait4 to collect zombies
		if _, err := unix.Wait4(-pid, &wstatus, unix.WALL, nil); err != unix.EINTR && err != nil {
			break
		}
	}
}

func setupChildForExternalDebugging(childPid int, forkRunner *forkexec.Runner) error {
	// Because we use the Ptrace option of the forkexec_runner,
	// we are now the debugger of the child.
	// We need to detach again to allow other debuggers.
	// But before that we continue execution from kill(SIGSTOP) until the exec-call.

	var wstatus unix.WaitStatus
	// child will SIGSTOP itself wait for that
	if _, err := unix.Wait4(childPid, &wstatus, 0, nil); err != nil {
		return err
	}
	if !wstatus.Stopped() || wstatus.StopSignal() != unix.SIGSTOP {
		return fmt.Errorf("unexpected wstatus: %d", wstatus)
	}
	// set TRACEEXEC to continue until exec
	if err := unix.PtraceSetOptions(childPid, unix.PTRACE_O_TRACEEXEC|unix.PTRACE_O_EXITKILL); err != nil {
		return err
	}
	// continue until exec
	if err := unix.PtraceCont(childPid, 0); err != nil {
		return err
	}
	// wait for ptrace trap from exec
	if _, err := unix.Wait4(childPid, &wstatus, 0, nil); err != nil {
		return err
	}
	if wstatus.StopSignal() != unix.SIGTRAP || wstatus.TrapCause() != unix.PTRACE_EVENT_EXEC {
		return fmt.Errorf("unexpected wstatus: %d", wstatus)
	}
	// SIGSTOP child again so it does not continue on detach
	if err := unix.Kill(childPid, unix.SIGSTOP); err != nil {
		return err
	}
	// detach so gdb can attach
	if err := unix.PtraceDetach(childPid); err != nil {
		return err
	}

	// write gdb command to stdout
	// use the random UID to find PID through the two "containers" (docker and our sandbox)
	// Because of the SIGSTOP from above gdb would (somehow) stop for that signal again,
	// so you'd need to use the "continue" command three times in gdb.
	// To prevent that use handle to not stop on that signal.
	// Then use the file command to load the symbols from the ELF we placed into the mount namespace.
	fmt.Printf("gdb -p `ps -o pid= -u %d` -ex \"handle SIGSTOP noprint\" -ex \"file target:/binary\"\n", forkRunner.Credential.Uid)
	return nil
}
