package check_service

import (
	"bytes"
	"checker/internal/algol60_printer"
	"checker/internal/utils"
	"fmt"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const executionCasesPath = "../compiler/src/test/resources/programs/"

func checkOneExecutionJunitCase(ip string) (checkerlib.Result, error) {
	var executionCases []string
	err := filepath.Walk(executionCasesPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".a60") {
				executionCases = append(executionCases, path)
			}
			return nil
		})
	if err != nil {
		log.Printf("Err reading testcases %s\n", err)
		return checkerlib.ResultInvalid, err
	}

	chosenCaseIdx := utils.CheckerRand.Int31n(int32(len(executionCases)))

	path := executionCases[chosenCaseIdx]
	filename := filepath.Base(path)

	code, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Err failed to read testcase[%s]: %s\n", filename, err)
		return checkerlib.ResultInvalid, err
	}

	// randomize code
	codeStr, _ := algol60_printer.RewriteCode(string(code), -1)

	expectedOutput, err := ioutil.ReadFile(strings.Replace(path, ".a60", ".outp", 1))
	if err != nil {
		log.Printf("Err failed to read testcase[%s]: %s\n", filename, err)
		return checkerlib.ResultInvalid, err
	}

	logReason := fmt.Sprintf("for testcase[%s]", filename)
	runResult, checkerResult, err := utils.CompileAndExec(codeStr, ip, logReason)
	if runResult == nil {
		return checkerResult, err
	}

	wstatus := unix.WaitStatus(runResult.Wstatus)
	if !strings.Contains(path, "illegal-index") {
		// Exit: 0
		if !strings.HasPrefix(runResult.Result, "Exit: ") || !wstatus.Exited() ||
			0 != wstatus.ExitStatus() {
			log.Printf("Unexpected result from running testcase[%s]: '%s'\n", filename, runResult.Result)
			return checkerlib.ResultFaulty, nil
		}
	} else {
		// Used illegal index - exit code 138
		if !strings.HasPrefix(runResult.Result, "Used illegal index") || !wstatus.Exited() ||
			138 != wstatus.ExitStatus() {
			log.Printf("Unexpected result from running testcase[%s]: '%s'\n", filename, runResult.Result)
			return checkerlib.ResultFaulty, nil
		}
	}

	// check stdout of execution
	if !bytes.Equal(expectedOutput, runResult.Stdout) {
		log.Printf("Wrong execution stdout for testcase[%s]\n", filename)
		return checkerlib.ResultFaulty, nil
	}

	log.Printf("Passed execution testcase[%s]\n", filename)

	return checkerlib.ResultOk, nil
}
