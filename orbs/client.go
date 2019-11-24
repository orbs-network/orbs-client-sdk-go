// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package orbs

import (
	"fmt"
	"net/http"

	"github.com/orbs-network/orbs-client-sdk-go/codec"
)

const PROTOCOL_VERSION = 1

type OrbsClient struct {
	Endpoint       string
	VirtualChainId uint32
	httpClient     *http.Client
	NetworkType    codec.NetworkType
}

func NewClient(endpoint string, virtualChainId uint32, networkType codec.NetworkType) *OrbsClient {
	// Customize the Transport to have larger connection pool
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100000
	defaultTransport.MaxIdleConnsPerHost = 100000

	turboClient := &http.Client{Transport: &defaultTransport}

	return &OrbsClient{
		Endpoint:       endpoint,
		httpClient:     turboClient,
		VirtualChainId: virtualChainId,
		NetworkType:    networkType,
	}
}
