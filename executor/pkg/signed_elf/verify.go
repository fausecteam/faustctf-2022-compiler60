package signed_elf

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"executor/pkg/compiler60_shared"
	"fmt"
	"golang.org/x/crypto/sha3"
	"time"
	"io/ioutil"
)

const signKeyPath = "/sign_key"
const defaultKey = "foobar"
const maxExpiryInFuture = 1 * time.Hour

func (signedElf SignedElf) VerifyStructure() error {
	now := time.Now()
	nowMilli := now.UnixNano() / 1e6
	endMilli := now.Add(maxExpiryInFuture).UnixNano() / 1e6
	if signedElf.Expiry < nowMilli || endMilli < signedElf.Expiry {
		return fmt.Errorf("signature expired")
	}

	if len(signedElf.Binary) > compiler60_shared.MaxBinarySize {
		return fmt.Errorf("binary too large")
	}

	if len(signedElf.Signature) != 256/8 {
		return fmt.Errorf("invalid signature size")
	}
	return nil
}

func getSignatureKey() (key []byte) {
	key, err := ioutil.ReadFile(signKeyPath)
	if err != nil {
		return []byte(defaultKey)
	}
	return
}

func (signedElf SignedElf) VerifyBinarySignature() error {
	if err := signedElf.VerifyStructure(); err != nil {
		return err
	}

	parsedElf, err := elf.NewFile(bytes.NewReader(signedElf.Binary))
	if err != nil {
		return fmt.Errorf("elf parser: %v", err)
	}

	hash := sha3.New256()

	hash.Write(getSignatureKey())

	// sign expiry
	var bs [8]byte
	binary.BigEndian.PutUint64(bs[:], uint64(signedElf.Expiry))
	hash.Write(bs[:])

	// sign binary except .rodata section (if any)
	rodataSection := parsedElf.Section(".rodata")
	if rodataSection == nil {
		hash.Write(signedElf.Binary)
	} else {
		start := rodataSection.Offset
		end := start + rodataSection.Size
		if int(start) > len(signedElf.Binary) || int(end) > len(signedElf.Binary) {
			return fmt.Errorf("invalid .rodata")
		}
		hash.Write(signedElf.Binary[:start])
		hash.Write(signedElf.Binary[end:])
	}

	digest := hash.Sum([]byte{})

	if bytes.Compare(digest, signedElf.Signature) != 0 {
		return fmt.Errorf("invalid signature")
	}
	return nil
}
