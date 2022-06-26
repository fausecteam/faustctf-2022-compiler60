package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"log"
)

// magic constant not exported by checkerlib
const FlagLookback = 5

// flag password is 160-bit input for sha1 to generate the filename
const flagPasswordRandomBytes = 20

type FlagWithPassword struct {
	Flag         string
	FlagPassword string
}

func (f FlagWithPassword) String() string {
	return fmt.Sprintf("{Flag:%s Pwd:%s}", f.Flag, f.FlagPassword)
}

const storeKeyFlagPasswords = "flag-passwords"

type flagPasswordsMap = map[int]string

func GetFilenameFromPassword(flagPassword string) string {
	// get sha1 hexdigest of flagPassword to get filename used by algol60-stdlib
	// and use that as public flag id
	bs := sha1.Sum([]byte(flagPassword))
	return hex.EncodeToString(bs[:])
}

func StorePasswordAndPublishFlagId(tick int, flagPassword string) {
	// publish flag id as filename with flag
	publicFlagId := GetFilenameFromPassword(flagPassword)
	checkerlib.SetFlagID(publicFlagId)

	// insert new flagPassword into map
	var storedPasswords flagPasswordsMap
	found := checkerlib.LoadState(storeKeyFlagPasswords, &storedPasswords)
	if !found {
		storedPasswords = make(map[int]string)
	}
	storedPasswords[tick] = flagPassword

	// delete storedPasswords older than 5 ticks
	for pwTick := range storedPasswords {
		if pwTick < tick-FlagLookback {
			delete(storedPasswords, pwTick)
		}
	}

	// store storedPasswords back
	checkerlib.StoreState(storeKeyFlagPasswords, storedPasswords)
}

func GenNewFlag(tick int) (r FlagWithPassword) {
	r.Flag = checkerlib.GetFlag(tick)

	var flagIdBs [flagPasswordRandomBytes]byte
	// fetch from crypto/rand
	_, _ = rand.Read(flagIdBs[:])
	// FlagPassword needs to be printable in source code so use base64
	// but we do not need padding
	r.FlagPassword = base64.RawStdEncoding.EncodeToString(flagIdBs[:])
	return
}

func GetFlagFromTick(tick int) (r FlagWithPassword, found bool) {
	r.Flag = checkerlib.GetFlag(tick)

	// load flagPassword from store
	var storedPasswords flagPasswordsMap
	found = checkerlib.LoadState(storeKeyFlagPasswords, &storedPasswords)
	if !found {
		log.Printf("loading flagPassword: '%s' not found in checker-state\n", storeKeyFlagPasswords)
		return
	}
	flagPassword, found := storedPasswords[tick]
	if !found {
		log.Printf("loading flagPassword: tick %d not in map of checker-state\n", tick)
		return
	}

	r.FlagPassword = flagPassword
	return
}
