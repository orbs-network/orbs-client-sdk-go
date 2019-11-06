// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/keys"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/pkg/errors"
	"time"
)

type RunQueryRequest struct {
	ProtocolVersion uint32
	VirtualChainId  uint32
	Timestamp       time.Time
	NetworkType     NetworkType
	PublicKey       []byte
	ContractName    string
	MethodName      string
	InputArguments  []interface{}
}

type RunQueryResponse struct {
	*ReadResponse
}

func EncodeRunQueryRequest(req *RunQueryRequest) ([]byte, error) {
	// validate
	if req.ProtocolVersion != 1 {
		return nil, errors.Errorf("expected ProtocolVersion 1, %d given", req.ProtocolVersion)
	}
	if len(req.PublicKey) != keys.ED25519_PUBLIC_KEY_SIZE_BYTES {
		return nil, errors.Errorf("expected PublicKey length %d, %d given", keys.ED25519_PUBLIC_KEY_SIZE_BYTES, len(req.PublicKey))
	}

	// encode method arguments
	inputArgumentArray, err := PackedArgumentsEncode(req.InputArguments)
	if err != nil {
		return nil, err
	}

	// encode network type
	networkType, err := networkTypeEncode(req.NetworkType)
	if err != nil {
		return nil, err
	}

	// encode request
	res := (&client.RunQueryRequestBuilder{
		SignedQuery: &protocol.SignedQueryBuilder{
			Query: &protocol.QueryBuilder{
				ProtocolVersion: primitives.ProtocolVersion(req.ProtocolVersion),
				VirtualChainId:  primitives.VirtualChainId(req.VirtualChainId),
				Timestamp:       primitives.TimestampNano(req.Timestamp.UnixNano()),
				Signer: &protocol.SignerBuilder{
					Scheme: protocol.SIGNER_SCHEME_EDDSA,
					Eddsa: &protocol.EdDSA01SignerBuilder{
						NetworkType:     networkType,
						SignerPublicKey: primitives.Ed25519PublicKey(req.PublicKey),
					},
				},
				ContractName:       primitives.ContractName(req.ContractName),
				MethodName:         primitives.MethodName(req.MethodName),
				InputArgumentArray: inputArgumentArray,
			},
		},
	}).Build()

	// return
	return res.Raw(), nil
}

func DecodeRunQueryResponse(buf []byte) (*RunQueryResponse, error) {
	// decode response
	res := client.RunQueryResponseReader(buf)
	if !res.IsValid() {
		return nil, errors.New("response is corrupt and cannot be decoded")
	}

	// return
	response, err := NewReadResponse(res, res.RawQueryResult(), res.QueryResult())
	if err != nil {
		return nil, err
	}
	return &RunQueryResponse{
		ReadResponse: response,
	}, nil
}
