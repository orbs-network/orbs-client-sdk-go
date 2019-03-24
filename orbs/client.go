// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package orbs

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
)

const PROTOCOL_VERSION = 1

type OrbsClient struct {
	Endpoint       string
	VirtualChainId uint32
	NetworkType    codec.NetworkType
}

func NewClient(endpoint string, virtualChainId uint32, networkType codec.NetworkType) *OrbsClient {
	return &OrbsClient{
		Endpoint:       endpoint,
		VirtualChainId: virtualChainId,
		NetworkType:    networkType,
	}
}
