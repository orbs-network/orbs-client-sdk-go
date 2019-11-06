// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"bytes"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/digest"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/hash"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/pkg/errors"
	"time"
)

type GetBlockRequest struct {
	ProtocolVersion uint32
	VirtualChainId  uint32
	BlockHeight     uint64
}

type GetBlockResponse struct {
	*Response
	TransactionsBlockHash   []byte
	TransactionsBlockHeader *TransactionsBlockHeader
	ResultsBlockHash        []byte
	ResultsBlockHeader      *ResultsBlockHeader
	Transactions            []*BlockTransaction
}

type TransactionsBlockHeader struct {
	ProtocolVersion uint32
	VirtualChainId  uint32
	BlockHeight     uint64
	PrevBlockHash   []byte
	Timestamp       time.Time
	NumTransactions uint32
}

type ResultsBlockHeader struct {
	ProtocolVersion        uint32
	VirtualChainId         uint32
	BlockHeight            uint64
	PrevBlockHash          []byte
	Timestamp              time.Time
	TransactionsBlockHash  []byte
	NumTransactionReceipts uint32
}

type BlockTransaction struct {
	TxId            []byte
	TxHash          []byte
	ProtocolVersion uint32
	VirtualChainId  uint32
	Timestamp       time.Time
	SignerPublicKey []byte
	ContractName    string
	MethodName      string
	InputArguments  []interface{}
	ExecutionResult ExecutionResult
	OutputArguments []interface{}
	OutputEvents    []*Event
}

func EncodeGetBlockRequest(req *GetBlockRequest) ([]byte, error) {
	// validate
	if req.ProtocolVersion != 1 {
		return nil, errors.Errorf("expected ProtocolVersion 1, %d given", req.ProtocolVersion)
	}

	// encode request
	res := (&client.GetBlockRequestBuilder{
		ProtocolVersion: primitives.ProtocolVersion(req.ProtocolVersion),
		VirtualChainId:  primitives.VirtualChainId(req.VirtualChainId),
		BlockHeight:     primitives.BlockHeight(req.BlockHeight),
	}).Build()

	// return
	return res.Raw(), nil
}

func DecodeGetBlockResponse(buf []byte) (*GetBlockResponse, error) {
	// decode response
	res := client.GetBlockResponseReader(buf)
	if !res.IsValid() {
		return nil, errors.New("response is corrupt and cannot be decoded")
	}

	// decode transactions
	var transactions []*BlockTransaction
	for txIterator := res.SignedTransactionsIterator(); txIterator.HasNext(); {
		tx := txIterator.NextSignedTransactions()

		// decode method arguments
		inputArgumentArray, err := protocol.PackedOutputArgumentsToNatives(tx.Transaction().RawInputArgumentArrayWithHeader())
		if err != nil {
			return nil, err
		}

		// add transaction
		transactions = append(transactions, &BlockTransaction{
			TxHash:          digest.CalcTxHash(tx.Transaction()),
			TxId:            digest.CalcTxId(tx.Transaction()),
			ProtocolVersion: uint32(tx.Transaction().ProtocolVersion()),
			VirtualChainId:  uint32(tx.Transaction().VirtualChainId()),
			Timestamp:       time.Unix(0, int64(tx.Transaction().Timestamp())),
			SignerPublicKey: tx.Transaction().Signer().Eddsa().SignerPublicKey(),
			ContractName:    string(tx.Transaction().ContractName()),
			MethodName:      string(tx.Transaction().MethodName()),
			InputArguments:  inputArgumentArray,
		})
	}

	// decode receipts
	for receiptIterator := res.TransactionReceiptsIterator(); receiptIterator.HasNext(); {
		receipt := receiptIterator.NextTransactionReceipts()
		for _, transaction := range transactions {
			if bytes.Equal(transaction.TxHash, receipt.Txhash()) {

				// decode execution result
				executionResult, err := executionResultDecode(receipt.ExecutionResult())
				if err != nil {
					return nil, err
				}
				transaction.ExecutionResult = executionResult

				// decode method arguments
				outputArgumentArray, err := protocol.PackedOutputArgumentsToNatives(receipt.RawOutputArgumentArrayWithHeader())
				if err != nil {
					return nil, err
				}
				transaction.OutputArguments = outputArgumentArray

				// decode events
				outputEventArray, err := PackedEventsDecode(receipt.RawOutputEventsArrayWithHeader())
				if err != nil {
					return nil, err
				}
				transaction.OutputEvents = outputEventArray

			}
		}
	}

	// return
	response, err := NewResponse(res)
	if err != nil {
		return nil, err
	}
	return &GetBlockResponse{
		Response:              response,
		TransactionsBlockHash: hash.CalcSha256(res.TransactionsBlockHeader().Raw()),
		TransactionsBlockHeader: &TransactionsBlockHeader{
			ProtocolVersion: uint32(res.TransactionsBlockHeader().ProtocolVersion()),
			VirtualChainId:  uint32(res.TransactionsBlockHeader().VirtualChainId()),
			BlockHeight:     uint64(res.TransactionsBlockHeader().BlockHeight()),
			PrevBlockHash:   res.TransactionsBlockHeader().PrevBlockHashPtr(),
			Timestamp:       time.Unix(0, int64(res.TransactionsBlockHeader().Timestamp())),
			NumTransactions: res.TransactionsBlockHeader().NumSignedTransactions(),
		},
		ResultsBlockHash: hash.CalcSha256(res.ResultsBlockHeader().Raw()),
		ResultsBlockHeader: &ResultsBlockHeader{
			ProtocolVersion:        uint32(res.ResultsBlockHeader().ProtocolVersion()),
			VirtualChainId:         uint32(res.ResultsBlockHeader().VirtualChainId()),
			BlockHeight:            uint64(res.ResultsBlockHeader().BlockHeight()),
			PrevBlockHash:          res.ResultsBlockHeader().PrevBlockHashPtr(),
			Timestamp:              time.Unix(0, int64(res.ResultsBlockHeader().Timestamp())),
			TransactionsBlockHash:  res.ResultsBlockHeader().TransactionsBlockHashPtr(),
			NumTransactionReceipts: res.ResultsBlockHeader().NumTransactionReceipts(),
		},
		Transactions: transactions,
	}, nil
}
