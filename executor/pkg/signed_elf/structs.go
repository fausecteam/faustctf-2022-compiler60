package signed_elf

import (
	"encoding/hex"
	"encoding/json"
	"executor/pkg/compiler60_shared"
)

// base64 is actually only 133% larger but keep some slack for the other fields
const MaxEncodedSize = 2 * compiler60_shared.MaxBinarySize

type HexByteSlice []byte

func (bs HexByteSlice) MarshalJSON() ([]byte, error) {
	hexString := hex.EncodeToString(bs)
	// quote to get a json string
	return json.Marshal(hexString)
}

func (bs *HexByteSlice) UnmarshalJSON(jsonBytes []byte) (err error) {
	var hexString string

	if err = json.Unmarshal(jsonBytes, &hexString); err != nil {
		return
	}
	*bs, err = hex.DecodeString(hexString)
	return
}

type SignedElf struct {
	Expiry    int64        `json:"expiry"`
	Binary    []byte       `json:"binary"`
	Signature HexByteSlice `json:"signature"`
}
