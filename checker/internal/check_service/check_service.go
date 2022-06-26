package check_service

import (
	"checker/internal/utils"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
)

func CheckService(ip string) (checkerlib.Result, error) {
	/*
		TODO more check service cases:
		- ELF contains stdlib in executable .text
		- Use Strings with newlines, emiojis, umlaute \xXX, \012 and \r\n\t and with '\\\"'
		- Use missleading names like binary base, syscall addr, gadget
		- possibly check, existence of symbols (stdlib and functions)
	*/

	// loop [4-5] times
	loopCount := int(4 + utils.CheckerRand.Int31n(2))
	for i := 0; i < loopCount; i++ {
		result, err := checkOneTestCase(ip)
		if err != nil || result != checkerlib.ResultOk {
			return result, err
		}
	}

	return checkerlib.ResultOk, nil
}

func checkOneTestCase(ip string) (checkerlib.Result, error) {
	rand := utils.CheckerRand.Int31n(10)
	if rand < 2 {
		// 2 / 10 chance
		return checkRodataRules(ip)
	}
	if rand < 7 {
		// 5 / 10 chance
		return checkOneExecutionJunitCase(ip)
	}
	// 3 / 10 chance
	return checkOneFailingJunitCase(ip)
}
