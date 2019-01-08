package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbsclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleTransfer(t *testing.T) {
	h := newHarness()
	defer h.shutdown()

	// create sender account
	sender, err := orbsclient.CreateAccount()
	require.NoError(t, err)

	// create receiver account
	receiver, err := orbsclient.CreateAccount()
	require.NoError(t, err)

	// create client
	endpoint := getEndpoint()
	client := orbsclient.NewOrbsClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	// create transfer transaction
	tx, txId, err := client.CreateTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		"BenchmarkToken",
		"transfer",
		uint64(10), receiver.AddressAsBytes())
	require.NoError(t, err)

	// send the transaction
	transferResponse, err := client.SendTransaction(tx)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, transferResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, transferResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, transferResponse.TransactionStatus)

	// check the transaction status
	statusResponse, err := client.GetTransactionStatus(txId)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, statusResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, statusResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, statusResponse.TransactionStatus)

	// check the transaction status receipt proof
	txProofResponse, err := client.GetTransactionReceiptProof(txId)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, txProofResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, txProofResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, txProofResponse.TransactionStatus)
	require.True(t, len(txProofResponse.PackedProof) > 20)

	// create balance query
	query, err := client.CreateQuery(
		receiver.PublicKey,
		"BenchmarkToken",
		"getBalance",
		receiver.AddressAsBytes())
	require.NoError(t, err)

	// send the query
	balanceResponse, err := client.SendQuery(query)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, balanceResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, balanceResponse.ExecutionResult)
	require.Equal(t, uint64(10), balanceResponse.OutputArguments[0])
}
