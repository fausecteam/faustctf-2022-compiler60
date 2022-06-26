package sandbox_runner

import (
	"encoding/json"
	"executor/pkg/compiler60_shared"
	"fmt"
	"github.com/criyle/go-sandbox/pkg/cgroup"
	"github.com/criyle/go-sandbox/runner"
	"golang.org/x/sys/unix"
	"os"
	"time"
)

type Runner struct {
	// binary as bytes to be executed
	ExecFile []byte

	cg cgroup.Cgroup
	// pipe fd connected to stdout of process
	outPipeRead    *os.File
	outPipeWriteFd int64

	execFileFd      *os.File
	execFileTmpPath string
}

type RunnerResult struct {
	Status  runner.Status
	WStatus unix.WaitStatus // exit status (signal number if signalled)
	Error   error           // potential detailed error message (for program runner error)

	Stdout []byte // output from stdout

	Time time.Duration // used user CPU time  (underlying type int64 in ns)
}

func (r RunnerResult) MarshalJSON() ([]byte, error) {
	var result string
	switch {
	case r.Status == runner.StatusTimeLimitExceeded:
		result = fmt.Sprintf("Time Limit Exceeded [%v]", r.Time)
	case r.Status == runner.StatusRunnerError:
		result = fmt.Sprintf("Runner error %s", r.Error)
	case r.Status == runner.StatusOutputLimitExceeded:
		result = fmt.Sprintf("Wrote too much on stdout [%v]", r.Time)
	case r.WStatus.Exited() && r.WStatus.ExitStatus() == 138:
		result = fmt.Sprintf("Used illegal index [%v]", r.Time)
	case r.WStatus.Exited():
		result = fmt.Sprintf("Exit: %d [%v]", r.WStatus.ExitStatus(), r.Time)
	case r.WStatus.Signaled():
		sig := r.WStatus.Signal()
		switch sig {
		case unix.SIGXCPU, unix.SIGKILL:
			result = fmt.Sprintf("Time Limit Exceeded [%v]", r.Time)
		case unix.SIGXFSZ:
			result = fmt.Sprintf("Wrote too large file (> %d)  [%v]", RlimitsFileSize, r.Time)
		case unix.SIGSYS:
			result = fmt.Sprintf("Illegal syscall [%v]", r.Time)
		default:
			result = fmt.Sprintf("Signaled: %s %v [%v]", unix.SignalName(sig), sig, r.Time)
		}
	}
	asJson := compiler60_shared.JsonRunnerResult{
		Result:     result,
		Stdout:     r.Stdout,
		UsedTimeMs: r.Time.Milliseconds(),
		Wstatus:    uint32(r.WStatus),
	}

	return json.Marshal(asJson)
}

func (r RunnerResult) String() string {
	marshalJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(marshalJSON)
}
