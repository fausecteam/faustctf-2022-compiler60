package utils

import (
	"checker/internal/api_client"
	"checker/internal/verify_elf"
	"executor/pkg/compiler60_shared"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"log"
)

func CompileAndExec(code string, ip string, reason string) (*compiler60_shared.JsonRunnerResult, checkerlib.Result, error) {
	log.Printf("Compiling code %s\n", reason)

	elf, compileErr, err := api_client.CompileCode(ip, code)
	if err != nil {
		return nil, checkerlib.ResultInvalid, err
	}
	// we do not expect any compile err
	if compileErr != nil {
		log.Printf("Compilation had unexpected error in line %d\n", compileErr.Line)
		return nil, checkerlib.ResultFaulty, nil
	}
	if elf == nil {
		return nil, checkerlib.ResultFaulty, nil
	}
	log.Printf("Compiled code %s (Elf size: %d, Expiry: %d)\n", reason, len(elf.Binary), elf.Expiry)

	if verify_elf.Verify(elf) == nil {
		return nil, checkerlib.ResultFaulty, nil
	}

	log.Printf("Executing code %s\n", reason)
	result, err := api_client.ExecuteCode(ip, elf)
	if err != nil {
		return nil, checkerlib.ResultInvalid, err
	}
	if result == nil {
		return nil, checkerlib.ResultFaulty, nil
	}
	log.Printf("Executed code %s (Wstatus: %d, Stdout len: %d)\n", reason, result.Wstatus, len(result.Stdout))

	return result, checkerlib.ResultInvalid, nil
}
