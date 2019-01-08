package orbs

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/encoding"
	"time"
)

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
