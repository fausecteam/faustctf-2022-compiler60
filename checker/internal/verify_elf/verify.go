package verify_elf

import (
	"bytes"
	"debug/elf"
	"executor/pkg/signed_elf"
	"log"
)

func Verify(signedElf *signed_elf.SignedElf) *elf.File {

	parsedElf, err := elf.NewFile(bytes.NewReader(signedElf.Binary))
	if err != nil {
		log.Printf("elf parser err %s\n", err)
		return nil
	}

	// TODO check stdlib?
	return parsedElf
}
