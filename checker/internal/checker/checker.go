package checker

import (
	"checker/internal/check_flag"
	"checker/internal/check_service"
	"checker/internal/place_flag"
	"checker/internal/utils"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
)

type Compiler60Checker struct {
}

func (c Compiler60Checker) PlaceFlag(ip string, team int, tick int) (checkerlib.Result, error) {

	utils.InitCheckerRand()

	result, err := place_flag.PlaceFlag(ip, tick)

	return utils.StoreSeedOnFailed(result, err)
}

func (c Compiler60Checker) CheckService(ip string, team int) (checkerlib.Result, error) {
	result, err := check_service.CheckService(ip)

	return utils.StoreSeedOnFailed(result, err)
}

func (c Compiler60Checker) CheckFlag(ip string, team int, tick int) (checkerlib.Result, error) {
	result, err := check_flag.CheckFlag(ip, tick)

	return utils.StoreSeedOnFailed(result, err)
}
