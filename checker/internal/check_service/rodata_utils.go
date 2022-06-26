package check_service

import (
	"bytes"
	"checker/internal/api_client"
	"checker/internal/verify_elf"
	"debug/elf"
	"encoding/binary"
	"executor/pkg/signed_elf"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"golang.org/x/sys/unix"
	"log"
	"strconv"
	"strings"
)

type extractedElf struct {
	signed_elf *signed_elf.SignedElf
	elf        *elf.File

	preRodata  []byte
	rodata     []byte
	postRodata []byte
}

func compileToElf(ip string, code string) (*extractedElf, checkerlib.Result, error) {
	log.Printf("Compiling code for checkRodataRules()\n")
	signedElf, compileErr, err := api_client.CompileCode(ip, code)
	if err != nil {
		return nil, checkerlib.ResultInvalid, err
	}
	// we do not expect any compile err
	if compileErr != nil {
		log.Printf("Compilation had unexpected error in line %d\n", compileErr.Line)
		return nil, checkerlib.ResultFaulty, nil
	}
	if signedElf == nil {
		return nil, checkerlib.ResultFaulty, nil
	}
	log.Printf("Compiled code for checkRodataRules() (Elf size: %d, Expiry: %d)\n", len(signedElf.Binary), signedElf.Expiry)

	return extractRodata(signedElf)
}

func extractRodata(signedElf *signed_elf.SignedElf) (*extractedElf, checkerlib.Result, error) {
	parsedElf := verify_elf.Verify(signedElf)
	if parsedElf == nil {
		return nil, checkerlib.ResultFaulty, nil
	}

	rodataSection := parsedElf.Section(".rodata")
	if rodataSection == nil {
		// we always expect .rodata in this function
		log.Println("Missing .rodata section")
		return nil, checkerlib.ResultFaulty, nil
	}
	start := rodataSection.Offset
	end := start + rodataSection.Size
	if int(start) > len(signedElf.Binary) || int(end) > len(signedElf.Binary) {
		log.Printf("invalid .rodata start:%d end:%d\n", start, end)
		return nil, checkerlib.ResultFaulty, nil
	}
	return &extractedElf{
		signed_elf: signedElf,
		elf:        parsedElf,
		preRodata:  signedElf.Binary[:start],
		rodata:     signedElf.Binary[start:end],
		postRodata: signedElf.Binary[end:],
	}, checkerlib.ResultInvalid, nil
}

func executeToConstant(ip string, signedElf *signed_elf.SignedElf, expectedResult int64) (checkerlib.Result, error) {
	log.Printf("Executing code for checkRodataRules()\n")
	runResult, err := api_client.ExecuteCode(ip, signedElf)
	if err != nil {
		return checkerlib.ResultInvalid, err
	}
	if runResult == nil {
		return checkerlib.ResultFaulty, nil
	}
	log.Printf("Executed code for checkRodataRules() (Wstatus: %d, Stdout len: %d)\n", runResult.Wstatus, len(runResult.Stdout))

	wstatus := unix.WaitStatus(runResult.Wstatus)
	// expect Exit: 0
	if !strings.HasPrefix(runResult.Result, "Exit: ") || !wstatus.Exited() ||
		0 != wstatus.ExitStatus() {
		log.Printf("Unexpected result from checkRodataRules(): '%s'\n", runResult.Result)
		return checkerlib.ResultFaulty, nil
	}

	// check stdout is equal to result
	expectedOutput := []byte(strconv.FormatInt(expectedResult, 10))
	if !bytes.Equal(expectedOutput, runResult.Stdout) {
		log.Printf("Wrong execution stdout for checkRodataRules() - expected %d\n", expectedResult)
		return checkerlib.ResultFaulty, nil
	}
	return checkerlib.ResultOk, nil
}

func (e extractedElf) checkRodataContainsInt64(value int64) bool {
	var valueBytes [8]byte
	binary.LittleEndian.PutUint64(valueBytes[:], uint64(value))
	return bytes.Contains(e.rodata, valueBytes[:])
}

func replaceConstInRodata(rodata []byte, orig, new int64) {
	var valueBytes [8]byte
	binary.LittleEndian.PutUint64(valueBytes[:], uint64(orig))
	idx := bytes.Index(rodata, valueBytes[:])
	binary.LittleEndian.PutUint64(valueBytes[:], uint64(new))
	copy(rodata[idx:idx+8], valueBytes[:])
}

func (e extractedElf) patchRodata(newRodata []byte) *signed_elf.SignedElf {
	if len(newRodata) != len(e.rodata) {
		log.Panicln("patchRodata(): rodata size changed")
	}
	var result signed_elf.SignedElf
	result = *e.signed_elf
	result.Binary = make([]byte, len(result.Binary))
	start := len(e.preRodata)
	end := start + len(newRodata)
	copy(result.Binary[:start], e.preRodata)
	copy(result.Binary[start:end], newRodata)
	copy(result.Binary[end:], e.postRodata)

	return &result
}
