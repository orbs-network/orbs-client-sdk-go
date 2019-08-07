// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package orbs

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/encoding"
	"io/ioutil"
	"path"
	"strings"
	"time"
)

type ProcessorType uint32
const PROCESSOR_TYPE_NATIVE = ProcessorType(1)
const PROCESSOR_TYPE_JAVASCRIPT = ProcessorType(2)

func (c *OrbsClient) CreateTransaction(publicKey []byte, privateKey []byte, contractName string, methodName string, inputArguments ...interface{}) (rawTransaction []byte, txId string, err error) {
	req, rawTxId, err := codec.EncodeSendTransactionRequest(&codec.SendTransactionRequest{
		ProtocolVersion: PROTOCOL_VERSION,
		VirtualChainId:  c.VirtualChainId,
		Timestamp:       time.Now(),
		NetworkType:     c.NetworkType,
		PublicKey:       publicKey,
		ContractName:    contractName,
		MethodName:      methodName,
		InputArguments:  inputArguments,
	}, privateKey)
	if err != nil {
		return nil, "", err
	}
	return req, encoding.EncodeHex(rawTxId), nil
}

func (c *OrbsClient) CreateQuery(publicKey []byte, contractName string, methodName string, inputArguments ...interface{}) (rawQuery []byte, err error) {
	return codec.EncodeRunQueryRequest(&codec.RunQueryRequest{
		ProtocolVersion: PROTOCOL_VERSION,
		VirtualChainId:  c.VirtualChainId,
		Timestamp:       time.Now(),
		NetworkType:     c.NetworkType,
		PublicKey:       publicKey,
		ContractName:    contractName,
		MethodName:      methodName,
		InputArguments:  inputArguments,
	})
}

func (c *OrbsClient) CreateDeployTransaction(publicKey []byte, privateKey []byte, contractName string, processorType ProcessorType, code ...[]byte) (rawTransaction []byte, txId string, err error) {
	params := []interface{} {
		contractName,
		uint32(processorType),
	}

	for _, c := range code {
		params = append(params, c)
	}

	return c.CreateTransaction(publicKey, privateKey, "_Deployments", "deployService", params...)
}

func (c *OrbsClient) createGetTransactionStatusPayload(txId string) (payload []byte, err error) {
	rawTxId, err := encoding.DecodeHex(txId)
	if err != nil {
		return nil, err
	}
	return codec.EncodeGetTransactionStatusRequest(&codec.GetTransactionStatusRequest{
		ProtocolVersion: PROTOCOL_VERSION,
		VirtualChainId:  c.VirtualChainId,
		TxId:            rawTxId,
	})
}

func (c *OrbsClient) createGetTransactionReceiptProofPayload(txId string) (payload []byte, err error) {
	rawTxId, err := encoding.DecodeHex(txId)
	if err != nil {
		return nil, err
	}
	return codec.EncodeGetTransactionReceiptProofRequest(&codec.GetTransactionReceiptProofRequest{
		ProtocolVersion: PROTOCOL_VERSION,
		VirtualChainId:  c.VirtualChainId,
		TxId:            rawTxId,
	})
}

func (c *OrbsClient) createGetBlockPayload(blockHeight uint64) (payload []byte, err error) {
	return codec.EncodeGetBlockRequest(&codec.GetBlockRequest{
		ProtocolVersion: PROTOCOL_VERSION,
		VirtualChainId:  c.VirtualChainId,
		BlockHeight:     blockHeight,
	})
}

// provides functionality required for CreateDeployTransaction
func ReadSourcesFromDir(dirname string) ([][]byte, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	var sources [][]byte
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") && !strings.HasSuffix(f.Name(), "_test.go") {
			source, err := ioutil.ReadFile(path.Join(dirname, f.Name()))
			if err != nil {
				return nil, err
			}

			sources = append(sources, source)
		}
	}

	return sources, nil
}