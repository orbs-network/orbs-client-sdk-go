package e2e

import (
	"fmt"
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
	endpoint := fmt.Sprintf("http://%s:%d", GAMMA_ENDPOINT, GAMMA_PORT)
	client := orbsclient.NewOrbsClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	// create transfer transaction payload
	payload, txId, err := client.CreateSendTransactionPayload(
		sender.PublicKey,
		sender.PrivateKey,
		"BenchmarkToken",
		"transfer",
		uint64(10), receiver.RawAddress)
	require.NoError(t, err)

	// send the payload
	transferResponse, err := client.SendTransaction(payload)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, transferResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, transferResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, transferResponse.TransactionStatus)

	// create get status payload
	payload, err = client.CreateGetTransactionStatusPayload(txId)
	require.NoError(t, err)

	// send the payload
	statusResponse, err := client.GetTransactionStatus(payload)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, statusResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, statusResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, statusResponse.TransactionStatus)

	// create balance method call payload
	payload, err = client.CreateCallMethodPayload(
		receiver.PublicKey,
		"BenchmarkToken",
		"getBalance",
		receiver.RawAddress)
	require.NoError(t, err)

	// send the payload
	balanceResponse, err := client.CallMethod(payload)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, balanceResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, balanceResponse.ExecutionResult)
	require.Equal(t, uint64(10), balanceResponse.OutputArguments[0])
}
