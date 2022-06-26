package sandbox_runner

import (
	"bytes"
	"github.com/criyle/go-sandbox/pkg/memfd"
	"golang.org/x/sys/unix"
	"os"
	"sync/atomic"
)

func (r *Runner) createMemFDOfBinary() (err error) {
	// create memfd, write bytes to it and seal
	r.execFileFd, err = memfd.DupToMemfd("binary", bytes.NewReader(r.ExecFile))
	return
}

func (r *Runner) createStdoutPipe() (err error) {
	var p [2]int
	err = unix.Pipe2(p[:], unix.O_CLOEXEC)
	if err != nil {
		return
	}
	pipeRead, pipeWrite := p[0], p[1]
	err = unix.SetNonblock(pipeRead, true)
	if err != nil {
		_ = unix.Close(pipeRead)
		_ = unix.Close(pipeWrite)
		return
	}

	r.outPipeRead = os.NewFile(uintptr(pipeRead), "stdoutRead")

	r.outPipeWriteFd = int64(pipeWrite)
	return
}

func (r *Runner) cleanupAtForkSync() {
	if r.execFileFd != nil {
		_ = r.execFileFd.Close() // ignore err
	}
	for {
		outPipeWriteFd := atomic.LoadInt64(&r.outPipeWriteFd)
		if outPipeWriteFd == 0 {
			break
		}
		if !atomic.CompareAndSwapInt64(&r.outPipeWriteFd, outPipeWriteFd, 0) {
			continue
		}
		_ = unix.Close(int(outPipeWriteFd)) // ignore err
	}
	if r.execFileTmpPath != "" {
		_ = os.Remove(r.execFileTmpPath) // ignore err
	}
}

func (r *Runner) cleanupRunner() {
	r.cleanupAtForkSync()
	if r.outPipeRead != nil {
		_ = r.outPipeRead.Close() // ignore err
	}
}
