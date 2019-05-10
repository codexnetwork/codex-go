package types

import (
	"encoding/hex"
	"encoding/json"
)

// Checksum256 a general type for xxx.Checksum256
type Checksum256 []byte

// MarshalJSON to json
func (t Checksum256) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(t))
}

// UnmarshalJSON from json
func (t *Checksum256) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = hex.DecodeString(s)
	return
}

func (t Checksum256) String() string {
	return hex.EncodeToString(t)
}
