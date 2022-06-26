package place_flag

import (
	"checker/internal/algol60_printer"
	"checker/internal/utils"
	"fmt"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"golang.org/x/sys/unix"
	"log"
	"strings"
)

func PlaceFlag(ip string, tick int) (checkerlib.Result, error) {
	flag := utils.GenNewFlag(tick)
	log.Printf("Place flag for tick %d: '%s'\n", tick, flag)

	code := placeFlagCode(flag)

	// randomize code
	code, _ = algol60_printer.RewriteCode(code, -1)

	runResult, checkerResult, err := utils.CompileAndExec(code, ip, "to place flag")
	if runResult == nil {
		return checkerResult, err
	}

	wstatus := unix.WaitStatus(runResult.Wstatus)
	if len(runResult.Stdout) != 0 {
		log.Printf("Unexpected stdout from place flag\n")
		return checkerlib.ResultFaulty, nil
	}
	if !strings.HasPrefix(runResult.Result, "Exit: ") || !wstatus.Exited() ||
		wstatus.ExitStatus() != 0 {
		log.Printf("Unexpected result from place flag: '%s'\n", runResult.Result)
		return checkerlib.ResultFaulty, nil
	}

	utils.StorePasswordAndPublishFlagId(tick, flag.FlagPassword)

	return checkerlib.ResultOk, nil
}

func placeFlagCode(flag utils.FlagWithPassword) (code string) {
	switch utils.CheckerRand.Int31n(5) {
	case 0:
		code = `
fd := openRW("%s");
'IF' fd < 0 'THEN'
  exit(1);
writestring(fd, "%s")`
		code = fmt.Sprintf(code, flag.FlagPassword, flag.Flag)

	case 1:
		code = `
'STRING' fl[%d];
'INTEGER' i;
fl := "%s";
fd := openRW("%s");
'IF' fd < 0 'THEN'
  exit(1);
i := -1;
'FOR' i := i + 1 'WHILE' fl[i] > 0 'DO'
  writechar(fd, fl[i])`
		code = fmt.Sprintf(code, len(flag.Flag)+1, flag.Flag, flag.FlagPassword)

	case 2:
		filename := utils.GetFilenameFromPassword(flag.FlagPassword)
		code = `
fd := openWOConfidential("%s");
'IF' fd < 0 'THEN'
  exit(1);
writestring(fd, "%s")
`
		code = fmt.Sprintf(code, filename, flag.Flag)

	case 3:
		filename := utils.GetFilenameFromPassword(flag.FlagPassword)
		code = `
'STRING' fl[1000];
'INTEGER' i;
fd := openWOConfidential("%s");
fl := "%s";
'IF' fd < 0 'THEN'
  exit(1);
i := -1;
'FOR' i := i + 1 'WHILE' fl[i] ≠ 0 'DO'
  writechar(fd, fl[i])`
		code = fmt.Sprintf(code, filename, flag.Flag)

	case 4:
		code = computeFlagCode(flag.Flag) +
			`fd := openRW("%s");
'IF' fd < 0 'THEN'
  exit(1);
writestring(fd, fl)`
		code = fmt.Sprintf(code, flag.FlagPassword)
	}

	code = `'BEGIN'
'INTEGER' fd;` + code + `
'END'`
	return
}

const flagChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_+/"

func computeFlagCode(flag string) string {
	addend := utils.CheckerRand.Int31n(401) - 200
	initialValue := make([]byte, len(flag))
	// create lines which assign the single bytes of the flag
	lines := make([]string, 0, len(flag))
	for i, b := range []byte(flag) {
		if utils.CheckerRand.Int31n(4) == 0 {
			initialValue[i] = b
		} else {
			initialValue[i] = flagChars[utils.CheckerRand.Int31n(int32(len(flagChars)))]
			x := int32(b) - addend
			line := fmt.Sprintf("fl[%d] := %d + x", i, x)
			lines = append(lines, line)
		}
	}
	// shuffle lines
	utils.CheckerRand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	linesStr := strings.Join(lines, ";\n  ")
	return fmt.Sprintf(`
'STRING' fl[%d];
'PROCEDURE' gFl(x);
  'INTEGER' x;
'BEGIN'
  fl := "%s";
  %s
'END';
gFl(%d);
`, len(flag)+1, initialValue, linesStr, addend)
}
