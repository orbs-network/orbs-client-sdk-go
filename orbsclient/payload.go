package orbsclient

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/base58"
	"time"
)

func (c *OrbsClient) CreateSendTransactionPayload(publicKey []byte, privateKey []byte, contractName string, methodName string, inputArguments ...interface{}) (payload []byte, txId string, err error) {
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
	return req, string(base58.Encode(rawTxId)), nil
}

func (c *OrbsClient) CreateCallMethodPayload(publicKey []byte, contractName string, methodName string, inputArguments ...interface{}) (payload []byte, err error) {
	return codec.EncodeCallMethodRequest(&codec.CallMethodRequest{
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

func (c *OrbsClient) CreateGetTransactionStatusPayload(txId string) (payload []byte, err error) {
	rawTxId, err := base58.Decode([]byte(txId))
	if err != nil {
		return nil, err
	}
	return codec.EncodeGetTransactionStatusRequest(&codec.GetTransactionStatusRequest{
		TxId: rawTxId,
	})
}
