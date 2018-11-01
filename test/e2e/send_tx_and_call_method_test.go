package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbsclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferFundsAndGetBalance(t *testing.T) {
	h := newHarness("localhost", 8080)
	defer h.shutdown()

	client := orbsclient.NewOrbsClient("http://localhost:8080", 42, codec.NETWORK_TYPE_TEST_NET)
	sender, err1 := orbsclient.CreateAccount()
	receiver, err2 := orbsclient.CreateAccount()

	require.NoError(t, err1)
	require.NoError(t, err2)

	transferTxPayload, txId, err := client.CreateSendTransactionPayload(
		sender.PublicKey,
		sender.PrivateKey,
		"BenchmarkToken",
		"transfer",
		uint64(10), receiver.RawAddress)

	require.NoError(t, err)

	transferTxResponse, err := client.SendTransaction(transferTxPayload)

	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, transferTxResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, transferTxResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, transferTxResponse.TransactionStatus)

	getStatusPayload, err := client.CreateGetTransactionStatusPayload(txId)
	require.NoError(t, err)

	getStatusResponse, err := client.GetTransactionStatus(getStatusPayload)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, getStatusResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, getStatusResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, getStatusResponse.TransactionStatus)

	getReceiverBalancePayload, err := client.CreateCallMethodPayload(
		receiver.PublicKey,
		"BenchmarkToken",
		"getBalance",
		receiver.RawAddress)
	getReceiverBalanceResponse, err := client.CallMethod(getReceiverBalancePayload)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, getReceiverBalanceResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, getReceiverBalanceResponse.ExecutionResult)
	require.Equal(t, uint64(10), getReceiverBalanceResponse.OutputArguments[0])
}
