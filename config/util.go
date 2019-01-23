package config

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"

	eos "github.com/eosforce/goeosforce"
)

// LoadJSONFile load a json file to obj
func LoadJSONFile(path string, obj interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, obj)
}

// ToSHA256Bytes from string to sha256
func ToSHA256Bytes(in string) (eos.SHA256Bytes, error) {
	if len(in) != 64 {
		return nil, errors.New("should be 64 hexadecimal characters")
	}

	bytes, err := hex.DecodeString(in)
	if err != nil {
		return nil, err
	}

	return eos.SHA256Bytes(bytes), nil
}
