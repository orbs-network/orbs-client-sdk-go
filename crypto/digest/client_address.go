// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

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
