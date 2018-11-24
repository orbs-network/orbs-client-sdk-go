package jsoncodec

import "encoding/json"

type Key struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    string // base58
}

func UnmarshalKeys(bytes []byte) (map[string]*Key, error) {
	keys := make(map[string]*Key)
	err := json.Unmarshal(bytes, &keys)
	return keys, err
}

func MarshalKeys(keys map[string]*Key) ([]byte, error) {
	return json.MarshalIndent(keys, "", "  ")
}
