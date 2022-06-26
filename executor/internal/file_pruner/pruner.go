package file_pruner

import (
	"executor/internal/sandbox_runner"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"os"
	"sync/atomic"
	"time"
)

func SchedulePruning() {
	go pruneDataDirectory()
}

const path = "/data"

// 7 ticks - flags have look back of 5
const deleteFilesAfter = 7 * 3 * time.Minute

// we do not get all dirents each prune execution (only direntsCountPerExecution)
// so we keep dirFDsKeptOpen to use in the next execution
const direntsCountPerExecution = 1024
const dirFDsKeptOpen = 3

var openDirectoriesChan = make(chan iteratingDir, dirFDsKeptOpen)

type iteratingDir struct {
	directoryFD    *os.File
	directoryFDNum int // == directoryFD.Fd()
}

func pruneDataDirectory() {
	var err error
	var dir iteratingDir
	// try to fetch already open directory fd
	select {
	case dir = <-openDirectoriesChan:
	default:
		// if channel empty open another one
		dir.directoryFD, err = os.Open(path)
		if err != nil {
			logSometimes("/data not found: %s\n", err)
			return
		}
		// will remove non-blocking - getdents() syscall is blocking anyway
		dir.directoryFDNum = int(dir.directoryFD.Fd())
	}

	// read the next batch of directory entries
	dirents, err := dir.directoryFD.Readdir(direntsCountPerExecution)
	if err != nil && err != io.EOF {
		logSometimes("/data readdir: %s\n", err)
		_ = dir.directoryFD.Close() // ignore err
		return
	}

	// loop over them and delete obvious ones directly
	expiry := time.Now().Add(-deleteFilesAfter)
	for _, info := range dirents {
		doDelete := false
		if info.IsDir() {
			// cannot delete directory (without something fancy like os.RemoveAll())
			logSometimes("Illegal directory in /data: '%s'\n", info.Name())
			continue
		}
		if info.ModTime().Before(expiry) {
			doDelete = true
		}
		if info.Mode() != sandbox_runner.AllowedFilePerm {
			doDelete = true
		}
		if info.Size() > sandbox_runner.RlimitsFileSize {
			doDelete = true
		}
		if doDelete {
			err = unix.Unlinkat(dir.directoryFDNum, info.Name(), 0)
			if err != nil {
				logSometimes("Cannot delete '%s' in /data: %s\n", info.Name(), err)
			}
			continue
		}
	}

	if len(dirents) == direntsCountPerExecution {
		// there might be more entries
		// try to hand over directory fd for one second to another execution
		select {
		case openDirectoriesChan <- dir:
			return
		case <-time.After(1 * time.Second):
		}
	}
	// else close fd
	_ = dir.directoryFD.Close() // ignore err
}

var lastLogTime int64 = 0
var logInterval = (2 * time.Minute).Milliseconds()

func logSometimes(format string, a ...any) {
	now := time.Now().UnixMilli()
	for {
		logTime := atomic.LoadInt64(&lastLogTime)
		if logTime+logInterval > now {
			return
		}
		if atomic.CompareAndSwapInt64(&lastLogTime, logTime, now) {
			fmt.Printf(format, a...)
			return
		}
	}
}
