package utils

import (
	"checker/internal/api_client"
	cryptoRand "crypto/rand"
	"encoding/binary"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"io"
	"log"
	"math/rand"
	"net/url"
)

// TODO easy way to set specific seed for CheckerRand for debugging

const storeKeyLastFailedSeed = "last-failed-seed"

// CheckerRand
// rand.Rand instance to be used for decisions made in the checker (like which specific checks to run)
// Seeded for each run except,
// if any check or placing flags fails the seed will be stored for the next run.
// This is to ensure teams fix problems this checker detects
// and do not rely on the problem occurring rarely.
// Ideally this would not be a global var, but a Compiler60Checker attribute,
// but I'm too lazy to pass the checker instance to each function.
var CheckerRand rand.Rand

var currentSeed uint32
var cachedStoredSeed uint32

var failedPreviously = false

func InitCheckerRand() {
	found := checkerlib.LoadState(storeKeyLastFailedSeed, &currentSeed)
	if !found {
		currentSeed = 0
	} else if currentSeed != 0 {
		log.Printf("CheckerRand: reusing stored seed %#x\n", currentSeed)
	}
	cachedStoredSeed = currentSeed

	for currentSeed == 0 {
		// use crypo/rand for the new seed
		var b [4]byte
		_, _ = cryptoRand.Read(b[:])
		currentSeed = binary.LittleEndian.Uint32(b[:])
		log.Printf("CheckerRand: using new seed %#x\n", currentSeed)
	}

	// NewSource internally only uses the int32 value
	source := rand.NewSource(int64(currentSeed))
	CheckerRand = *rand.New(source)
}

func StoreSeedOnFailed(result checkerlib.Result, err error) (checkerlib.Result, error) {
	// map EOF to DOWN
	if urlErr, ok := err.(*url.Error); ok {
		if urlErr.Err == io.EOF {
			result = checkerlib.ResultDown
			err = nil
		}
	}
	// this is a bit ugly cause 503 could not only come from HAProxy but also executor if overloaded
	// map 503 errors to DOWN
	if err == api_client.HTTP503Error {
		result = checkerlib.ResultDown
		err = nil
	}

	valueToStore := currentSeed

	failed := err != nil || result != checkerlib.ResultOk
	failedPreviously = failedPreviously || failed
	if !failedPreviously {
		// when not failed we still need to store 0
		// to avoid using the seed again when 'problem' disappears on the checked service
		valueToStore = 0
	}

	if cachedStoredSeed != valueToStore {
		log.Printf("CheckerRand: storing seed %#x\n", valueToStore)
		checkerlib.StoreState(storeKeyLastFailedSeed, valueToStore)
		cachedStoredSeed = valueToStore
	}

	return result, err
}
