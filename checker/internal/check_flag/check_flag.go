package check_flag

import (
	"checker/internal/algol60_printer"
	"checker/internal/utils"
	"fmt"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"strings"
)

func CheckFlag(ip string, tick int) (checkerlib.Result, error) {
	flag, found := utils.GetFlagFromTick(tick)
	if !found {
		return checkerlib.ResultFlagNotFound, nil
	}
	log.Printf("Check flag for tick %d: '%s'\n", tick, flag)

	code, expectedStdout := checkFlagCode(flag)

	// randomize code
	code, _ = algol60_printer.RewriteCode(code, -1)

	runResult, checkerResult, err := utils.CompileAndExec(code, ip, "to check flag")
	if runResult == nil {
		return checkerResult, err
	}

	wstatus := unix.WaitStatus(runResult.Wstatus)
	if !strings.HasPrefix(runResult.Result, "Exit: ") || !wstatus.Exited() ||
		(wstatus.ExitStatus() != 0 && wstatus.ExitStatus() != 1) {
		log.Printf("Unexpected result from check flag: '%s'\n", runResult.Result)
		return checkerlib.ResultFaulty, nil
	}

	if wstatus.ExitStatus() == 1 {
		log.Printf("Flag not found: '%s' Stdout len: %d\n", runResult.Result, len(runResult.Stdout))
		return checkerlib.ResultFlagNotFound, nil
	}
	stdout := string(runResult.Stdout)
	if expectedStdout != stdout {
		if len(stdout) > 50 {
			log.Printf("Invalid Flag: '%s' Stdout len: %d\n", runResult.Result, len(runResult.Stdout))
		} else {
			log.Printf("Invalid Flag: '%s' Expected-Actual: '%s'!='%s'\n", runResult.Result, expectedStdout, stdout)
		}
		return checkerlib.ResultFlagNotFound, nil
	}

	return checkerlib.ResultOk, nil
}

func checkFlagCode(flag utils.FlagWithPassword) (code string, stdout string) {
	stdout = flag.Flag
	switch utils.CheckerRand.Int31n(4) {
	case 0:
		code = `
fd := openRO("%s");
'IF' fd < 0 'THEN'
  exit(1);
outstring(readstring(fd))`
		code = fmt.Sprintf(code, flag.FlagPassword)

	case 1:
		addend := utils.CheckerRand.Int31n(401) - 200
		code = `
'STRING' pw[%d];
pw := "%s";
fd := openRO(pw);
'IF' fd < 0 'THEN'
  exit(1);
fd := fd - (%+d);
outstring(readstring(fd %+d))`
		code = fmt.Sprintf(code, len(flag.FlagPassword)+1, flag.FlagPassword, addend, addend)

	case 2:
		// rotate string bytes / ROT256
		addend := utils.CheckerRand.Int31n(401) - 200
		code = `
'INTEGER' c;
fd := openRO("%s");
'IF' fd < 0 'THEN'
  exit(1);
'FOR' c := readchar(fd) 'WHILE' c ≥ 0 'DO'
  outchar(c %+d);`
		code = fmt.Sprintf(code, flag.FlagPassword, addend)
		stdoutRotate := make([]byte, len(flag.Flag))
		for i := 0; i < len(flag.Flag); i++ {
			stdoutRotate[i] = uint8(int32(flag.Flag[i]) + addend)
		}
		stdout = string(stdoutRotate)

	case 3:
		// print flag as ascii numbers
		code = `
'INTEGER' c;
fd := openRO("%s");
'IF' fd < 0 'THEN'
  exit(1);
'FOR' c := readchar(fd) 'WHILE' c ≠ -1 'DO' 'BEGIN'
  outinteger(c);
  outchar(10);
'END'`
		code = fmt.Sprintf(code, flag.FlagPassword)
		stdout = ""
		for _, b := range []byte(flag.Flag) {
			stdout += strconv.FormatInt(int64(b), 10)
			stdout += "\n"
		}
	}

	code = `'BEGIN'
'INTEGER' fd;` + code + `
'END'
`
	return
}
