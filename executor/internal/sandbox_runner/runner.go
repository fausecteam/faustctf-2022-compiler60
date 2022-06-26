package sandbox_runner

import (
	"context"
	"github.com/criyle/go-sandbox/pkg/cgroup"
	"github.com/criyle/go-sandbox/runner"
	"os"
	"sync"
)

// for each of the possible parallel runner store its
var parallelRunners chan *cgroup.Cgroup
var parallelRunnersOnce sync.Once

func (r *Runner) Run(c context.Context) (result RunnerResult) {
	parallelRunnersOnce.Do(func() {
		parallelRunners = make(chan *cgroup.Cgroup, MaxConcurrentRunners)
		for i := 0; i < MaxConcurrentRunners; i++ {
			parallelRunners <- nil
		}
	})

	// limit parallel runners
	// try to catch a cgroup or timeout via context
	select {
	case <-c.Done():
		// result will not be used so keep it empty
		return
	case cg := <-parallelRunners:
		if cg != nil {
			r.cg = *cg
		} else {
			var err error
			cg, err := configureNewCgroup()
			if err != nil {
				result.Error = err
				result.Status = runner.StatusRunnerError
				return
			}
			r.cg = cg
		}
		// return cgroup
		defer func() { parallelRunners <- &r.cg }()
	}

	// max sure to close all fds (like memfd and pipe)
	defer r.cleanupRunner()

	forkRunner, err := r.setupRunner()
	if err != nil {
		result.Error = err
		result.Status = runner.StatusRunnerError
		return
	}

	// kill process after MaxRuntime
	ctx, cancel := context.WithTimeout(c, MaxRuntime)
	if IsDebug() {
		// disable timeout for debugging
		ctx, cancel = context.WithCancel(c)
	}
	defer cancel()

	result, err = r.startAndWait(ctx, forkRunner)
	if err != nil {
		result.Error = err
		result.Status = runner.StatusRunnerError
	}
	return
}

func IsDebug() bool {
	val, found := os.LookupEnv("DEBUG")
	return found && "1" == val
}
