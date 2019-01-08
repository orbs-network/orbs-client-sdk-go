package _manual

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbsclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferFundsAndGetBalance(t *testing.T) {
	client := orbsclient.NewOrbsClient("http://us-east-1.global.nodes.staging.orbs-test.com:80", 42, codec.NETWORK_TYPE_TEST_NET)
	sender, err1 := orbsclient.CreateAccount()
	receiver, err2 := orbsclient.CreateAccount()

	require.NoError(t, err1)
	require.NoError(t, err2)

	transferTxPayload, txId, err := client.CreateTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		"BenchmarkToken",
		"transfer",
		uint64(10), receiver.RawAddress)
	require.NoError(t, err)

	transferTxResponse, err := client.SendTransaction(transferTxPayload)

	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, transferTxResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, transferTxResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, transferTxResponse.TransactionStatus)
	require.Equal(t, "transfer of 10 failed since balance is only 0", transferTxResponse.OutputArguments[0])

	getStatusPayload, err := client.CreateGetTransactionStatusPayload(txId)
	require.NoError(t, err)
	getStatusResponse, err := client.GetTransactionStatus(getStatusPayload)
	require.NoError(t, err)

	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, getStatusResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, getStatusResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, getStatusResponse.TransactionStatus)

	getReceiverBalancePayload, err := client.CreateRunQueryPayload(
		sender.PublicKey,
		"BenchmarkToken",
		"getBalance",
		sender.RawAddress)
	getReceiverBalanceResponse, err := client.RunQuery(getReceiverBalancePayload)

	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, getReceiverBalanceResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, getReceiverBalanceResponse.ExecutionResult)
	require.Equal(t, uint64(0), getReceiverBalanceResponse.OutputArguments[0])
}
