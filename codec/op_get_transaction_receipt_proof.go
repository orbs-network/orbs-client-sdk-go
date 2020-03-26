// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"github.com/orbs-network/crypto-lib-go/crypto/digest"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/pkg/errors"
)

type GetTransactionReceiptProofRequest struct {
	ProtocolVersion uint32
	VirtualChainId  uint32
	TxId            []byte
}

type GetTransactionReceiptProofResponse struct {
	*TransactionResponse
	PackedProof   []byte
	PackedReceipt []byte
}

func EncodeGetTransactionReceiptProofRequest(req *GetTransactionReceiptProofRequest) ([]byte, error) {
	// validate
	if req.ProtocolVersion != 1 {
		return nil, errors.Errorf("expected ProtocolVersion 1, %d given", req.ProtocolVersion)
	}
	if len(req.TxId) != digest.TX_ID_SIZE_BYTES {
		return nil, errors.Errorf("expected TxId length %d, %d given", digest.TX_ID_SIZE_BYTES, len(req.TxId))
	}

	// extract txid
	txHash, txTimestamp, err := digest.ExtractTxId(req.TxId)
	if err != nil {
		return nil, err
	}

	// encode request
	res := (&client.GetTransactionReceiptProofRequestBuilder{
		TransactionRef: &client.TransactionRefBuilder{
			ProtocolVersion:      primitives.ProtocolVersion(req.ProtocolVersion),
			VirtualChainId:       primitives.VirtualChainId(req.VirtualChainId),
			TransactionTimestamp: txTimestamp,
			Txhash:               primitives.Sha256(txHash),
		},
	}).Build()

	// return
	return res.Raw(), nil
}

func DecodeGetTransactionReceiptProofResponse(buf []byte) (*GetTransactionReceiptProofResponse, error) {
	// decode response
	res := client.GetTransactionReceiptProofResponseReader(buf)
	if !res.IsValid() {
		return nil, errors.New("response is corrupt and cannot be decoded")
	}

	txResponse, err := NewTransactionResponse(res)
	if err != nil {
		return nil, err
	}

	return &GetTransactionReceiptProofResponse{
		TransactionResponse: txResponse,
		PackedProof:         res.PackedProof(),
		PackedReceipt:       res.TransactionReceipt().Raw(),
	}, nil
}
