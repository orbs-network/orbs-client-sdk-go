package orbsclient

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
)

const PROTOCOL_VERSION = 1

type OrbsClient struct {
	Endpoint       string
	VirtualChainId uint32
	NetworkType    codec.NetworkType
}

func NewOrbsClient(endpoint string, virtualChainId uint32, networkType codec.NetworkType) *OrbsClient {
	return &OrbsClient{
		Endpoint:       endpoint,
		VirtualChainId: virtualChainId,
		NetworkType:    networkType,
	}
}
