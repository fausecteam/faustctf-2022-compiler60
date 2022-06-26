package check_service

import (
	"bytes"
	"checker/internal/algol60_printer"
	"checker/internal/utils"
	"executor/pkg/signed_elf"
	"fmt"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"log"
)

const (
	opAdd = iota
	opSub
	opMul
)

func checkRodataRules(ip string) (checkerlib.Result, error) {
	template, op := checkRodataTemplate()

	// COMPILE 1
	const1_1, const1_2 := randConstant(), randConstant()
	code := fmt.Sprintf(template, const1_1, const1_2)
	code = algol60_printer.RewriteCodeDeterministic(code)
	elf1, checkerResult, err := compileToElf(ip, code)
	if elf1 == nil {
		return checkerResult, err
	}

	// COMPILE 2
	const2_1, const2_2 := randConstant(), randConstant()
	code = fmt.Sprintf(template, const2_1, const2_2)
	code = algol60_printer.RewriteCodeDeterministic(code)
	elf2, checkerResult, err := compileToElf(ip, code)
	if elf2 == nil {
		return checkerResult, err
	}

	// check everything except .rodata equal
	if !bytes.Equal(elf1.preRodata, elf2.preRodata) || !bytes.Equal(elf1.postRodata, elf2.postRodata) {
		log.Printf("binary except .rodata not equal after changed constant")
		return checkerlib.ResultFaulty, nil
	}
	// constant values must be in .rodata
	if !elf1.checkRodataContainsInt64(const1_1) ||
		!elf1.checkRodataContainsInt64(const1_2) ||
		!elf2.checkRodataContainsInt64(const2_1) ||
		!elf2.checkRodataContainsInt64(const2_2) {
		log.Println(".rodata did not contain constant")
		return checkerlib.ResultFaulty, nil
	}

	const3_1, const3_2 := randConstant(), randConstant()
	if const1_1 == const1_2 {
		// edge case: order in .rodata is not fixed so make them also equal
		const3_2 = const3_1
	}

	var result1, result2, result3 int64
	switch op {
	case opAdd:
		result1 = const1_1 + const1_2
		result2 = const2_1 + const2_2
		result3 = const3_1 + const3_2
	case opSub:
		result1 = const1_1 - const1_2
		result2 = const2_1 - const2_2
		result3 = const3_1 - const3_2
	case opMul:
		result1 = const1_1 * const1_2
		result2 = const2_1 * const2_2
		result3 = const3_1 * const3_2
	}

	if utils.CheckerRand.Int31n(2) != 0 {
		log.Println("swap .rodata of elfs")
		elf1.signed_elf = elf2.patchRodata(elf1.rodata)
		elf2.signed_elf = elf1.patchRodata(elf2.rodata)
	}

	// EXECUTE 1
	checkerResult, err = executeToConstant(ip, elf1.signed_elf, result1)
	if checkerResult != checkerlib.ResultOk {
		return checkerResult, err
	}

	// EXECUTE 2
	checkerResult, err = executeToConstant(ip, elf2.signed_elf, result2)
	if checkerResult != checkerlib.ResultOk {
		return checkerResult, err
	}

	// EXECUTE patched elf with new constants
	patchedRodata := make([]byte, len(elf1.rodata))
	copy(patchedRodata, elf1.rodata)
	replaceConstInRodata(patchedRodata, const1_1, const3_1)
	replaceConstInRodata(patchedRodata, const1_2, const3_2)
	var patchedElf *signed_elf.SignedElf
	if utils.CheckerRand.Int31n(2) != 0 {
		patchedElf = elf1.patchRodata(patchedRodata)
	} else {
		patchedElf = elf2.patchRodata(patchedRodata)
	}
	checkerResult, err = executeToConstant(ip, patchedElf, result3)
	if checkerResult != checkerlib.ResultOk {
		return checkerResult, err
	}

	log.Println("Passed checkRodataRules()")

	return checkerlib.ResultOk, nil
}

func checkRodataTemplate() (code string, op int32) {
	op = utils.CheckerRand.Int31n(3)
	opStr := []string{"+", "-", "×"}[op]
	switch utils.CheckerRand.Int31n(5) {
	case 0:
		code = `outinteger(%%d %s %%d)`
		code = fmt.Sprintf(code, opStr)

	case 1:
		if op == opMul {
			code = `outinteger(-%%d %s (-%%d))`
		} else {
			code = `outinteger(-(-%%d %s (-%%d)))`
		}
		code = fmt.Sprintf(code, opStr)

	case 2:
		code = `'INTEGER' aa,bb;
aa := %%d;
bb := %%d;
outinteger(aa %s bb)`
		code = fmt.Sprintf(code, opStr)

	case 3:
		code = `'STRING' intstr[21];
'INTEGER' i;
intstr := integer2string(%%d %s %%d);
i := -1;
'FOR' i := i+1 'WHILE' intstr[i] > 0 'DO'
  outchar(intstr[i])`
		code = fmt.Sprintf(code, opStr)

	case 4:
		code = `'PROCEDURE' main(a,b);
  'INTEGER' a;
  'INTEGER' b;
  outinteger(a %s b);
main(%%d, %%d)`
		code = fmt.Sprintf(code, opStr)
	}

	code = `'BEGIN'
` + code + `
'END'
`
	return
}

func randConstant() (r int64) {
	// avoid 0 and 1 because they are not treated as constants by compiler
	// use only positive value because grammar does not support them (-2 is 0-2)
	for ; r == 0 || r == 1; r = utils.CheckerRand.Int63() {
	}
	return
}
