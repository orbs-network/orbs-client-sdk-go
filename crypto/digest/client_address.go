package digest

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/hash"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/keys"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/pkg/errors"
)

const (
	CLIENT_ADDRESS_SIZE_BYTES    = 20
	CLIENT_ADDRESS_SHA256_OFFSET = hash.SHA256_HASH_SIZE_BYTES - CLIENT_ADDRESS_SIZE_BYTES
)

func CalcClientAddressOfEd25519PublicKey(publicKey primitives.Ed25519PublicKey) (primitives.ClientAddress, error) {
	if len(publicKey) != keys.ED25519_PUBLIC_KEY_SIZE_BYTES {
		return nil, errors.New("transaction is not signed by a valid Signer")
	}
	res := hash.CalcSha256(publicKey)[CLIENT_ADDRESS_SHA256_OFFSET:]
	return primitives.ClientAddress(res), nil
}
