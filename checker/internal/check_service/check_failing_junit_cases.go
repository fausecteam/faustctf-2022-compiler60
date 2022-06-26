package check_service

import (
	"checker/internal/algol60_printer"
	"checker/internal/api_client"
	"checker/internal/utils"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const failingNamesCasesPath = "../compiler/src/test/resources/names/"
const failingTypesCasesPath = "../compiler/src/test/resources/types/"

func checkOneFailingJunitCase(ip string) (checkerlib.Result, error) {
	namesCases, err := ioutil.ReadDir(failingNamesCasesPath)
	if err != nil {
		log.Printf("Err reading testcases %s\n", err)
		return checkerlib.ResultInvalid, err
	}

	typesCases, err := ioutil.ReadDir(failingTypesCasesPath)
	if err != nil {
		log.Printf("Err reading testcases %s\n", err)
		return checkerlib.ResultInvalid, err
	}

	namesCount := int32(len(namesCases))
	testCases := append(namesCases, typesCases...)

	chosenCaseIdx := utils.CheckerRand.Int31n(int32(len(testCases)))

	filename := testCases[chosenCaseIdx].Name()
	path := failingNamesCasesPath + filename
	if chosenCaseIdx >= namesCount {
		path = failingTypesCasesPath + filename
	}

	expectedErrorLineStr := filename[strings.Index(filename, "-L")+2 : len(filename)-len(".a60")]
	expectedErrorLine, err := strconv.Atoi(expectedErrorLineStr)
	if err != nil {
		log.Printf("Err parsing line in testcase filename[%s]: %s\n", filename, err)
		return checkerlib.ResultInvalid, err
	}

	code, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Err failed to read testcase[%s]: %s\n", filename, err)
		return checkerlib.ResultInvalid, err
	}

	// randomize code
	codeStr, expectedErrorLine := algol60_printer.RewriteCode(string(code), expectedErrorLine)

	log.Printf("Compiling failing testcase[%s]\n", filename)

	elf, compileErr, err := api_client.CompileCode(ip, codeStr)
	if err != nil {
		return checkerlib.ResultInvalid, err
	}
	// we never expect a successfully compiled elf
	if elf != nil {
		log.Printf("Successful compile for failing testcase[%s]\n", filename)
		return checkerlib.ResultFaulty, nil
	}
	if compileErr == nil {
		return checkerlib.ResultFaulty, nil
	}

	// check line of error
	if expectedErrorLine != compileErr.Line {
		log.Printf("Wrong error line for testcase[%s]: %d\n", filename, compileErr.Line)
		return checkerlib.ResultFaulty, nil
	}

	log.Printf("Passed failing testcase[%s]\n", filename)

	return checkerlib.ResultOk, nil
}
