// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/digest"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/pkg/errors"
	"time"
)

type GetTransactionStatusRequest struct {
	ProtocolVersion uint32
	VirtualChainId  uint32
	TxId            []byte
}

type GetTransactionStatusResponse struct {
	RequestStatus     RequestStatus
	TxHash            []byte
	ExecutionResult   ExecutionResult
	OutputArguments   []interface{}
	OutputEvents      []*Event
	TransactionStatus TransactionStatus
	BlockHeight       uint64
	BlockTimestamp    time.Time
}

func EncodeGetTransactionStatusRequest(req *GetTransactionStatusRequest) ([]byte, error) {
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
	res := (&client.GetTransactionStatusRequestBuilder{
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

func DecodeGetTransactionStatusResponse(buf []byte) (*GetTransactionStatusResponse, error) {
	// decode response
	res := client.GetTransactionStatusResponseReader(buf)
	if !res.IsValid() {
		return nil, errors.New("response is corrupt and cannot be decoded")
	}

	// decode request status
	requestStatus, err := requestStatusDecode(res.RequestResult().RequestStatus())
	if err != nil {
		return nil, err
	}

	// decode execution result
	executionResult := EXECUTION_RESULT_NOT_EXECUTED
	if len(res.RawTransactionReceipt()) > 0 {
		executionResult, err = executionResultDecode(res.TransactionReceipt().ExecutionResult())
		if err != nil {
			return nil, err
		}
	}

	// decode method arguments
	outputArgumentArray, err := protocol.PackedOutputArgumentsToNatives(res.TransactionReceipt().RawOutputArgumentArrayWithHeader())
	if err != nil {
		return nil, err
	}

	// decode events
	outputEventArray, err := PackedEventsDecode(res.TransactionReceipt().RawOutputEventsArrayWithHeader())
	if err != nil {
		return nil, err
	}

	// decode transaction status
	transactionStatus, err := transactionStatusDecode(res.TransactionStatus())
	if err != nil {
		return nil, err
	}

	// return
	return &GetTransactionStatusResponse{
		RequestStatus:     requestStatus,
		TxHash:            res.TransactionReceipt().Txhash(),
		ExecutionResult:   executionResult,
		OutputArguments:   outputArgumentArray,
		OutputEvents:      outputEventArray,
		TransactionStatus: transactionStatus,
		BlockHeight:       uint64(res.RequestResult().BlockHeight()),
		BlockTimestamp:    time.Unix(0, int64(res.RequestResult().BlockTimestamp())),
	}, nil
}
