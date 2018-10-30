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
