package orbsclient

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/base58"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/digest"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/keys"
)

type OrbsAccount struct {
	PublicKey  []byte
	PrivateKey []byte
	Address    string
	RawAddress []byte
}

func CreateAccount() (*OrbsAccount, error) {
	keyPair, err := keys.GenerateEd25519Key()
	if err != nil {
		return nil, err
	}

	rawAddress, err := digest.CalcClientAddressOfEd25519PublicKey(keyPair.PublicKey())
	if err != nil {
		return nil, err
	}

	return &OrbsAccount{
		PublicKey:  keyPair.PublicKey(),
		PrivateKey: keyPair.PrivateKey(),
		Address:    string(base58.Encode(rawAddress)),
		RawAddress: rawAddress,
	}, nil
}
