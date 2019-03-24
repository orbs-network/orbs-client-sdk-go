// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

type NetworkType string

const (
	NETWORK_TYPE_MAIN_NET NetworkType = "MAIN_NET"
	NETWORK_TYPE_TEST_NET NetworkType = "TEST_NET"
)

func (x NetworkType) String() string {
	return string(x)
}

func networkTypeEncode(networkType NetworkType) (protocol.SignerNetworkType, error) {
	switch networkType {
	case NETWORK_TYPE_MAIN_NET:
		return protocol.NETWORK_TYPE_MAIN_NET, nil
	case NETWORK_TYPE_TEST_NET:
		return protocol.NETWORK_TYPE_TEST_NET, nil
	default:
		return protocol.NETWORK_TYPE_RESERVED, errors.Errorf("unsupported network type given %s", networkType)
	}
}
