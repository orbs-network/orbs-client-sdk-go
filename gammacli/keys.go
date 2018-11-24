package main

import (
	"encoding/json"
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/gammacli/jsoncodec"
	"github.com/orbs-network/orbs-client-sdk-go/orbsclient"
	"io/ioutil"
)

func commandGenerateTestKeys() {
	keys := make(map[string]*jsoncodec.Key)
	for i := 0; i < 10; i++ {
		account, err := orbsclient.CreateAccount()
		if err != nil {
			die("could not create Orbs account")
		}
		user := fmt.Sprintf("user%d", i+1)
		keys[user] = &jsoncodec.Key{
			PrivateKey: account.PrivateKey,
			PublicKey:  account.PublicKey,
			Address:    account.Address,
		}
	}

	bytes, err := jsoncodec.MarshalKeys(keys)
	if err != nil {
		die("could not encode keys to json\n\n%s", err.Error())
	}

	filename := *flagKeyFile
	if filename == "" {
		filename = TEST_KEYS_FILENAME
	}
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		die("could not write keys to file\n\n%s", err.Error())
	}

	if !doesFileExist(filename) {
		die("file not found after write")
	}

	log("10 new test keys written successfully to '%s'\n", filename)
}

func getTestKeyFromFile(id string) *jsoncodec.Key {
	bytes, err := ioutil.ReadFile(*flagKeyFile)
	if err != nil {
		die("could not open keys file\n\n%s", err.Error())
	}

	keys, err := jsoncodec.UnmarshalKeys(bytes)
	err = json.Unmarshal(bytes, &keys)
	if err != nil {
		die("failed parsing keys json file\n\n%s", *flagKeyFile, err.Error())
	}

	key, found := keys[id]
	if !found {
		die("key with id '%s' not found in key file '%s'", id, *flagKeyFile)
	}

	return key
}
