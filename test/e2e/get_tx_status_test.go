package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTxStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping E2E tests in short mode")
	}

	h := newHarness()
	defer h.shutdown()

	// create sender account
	sender, err := orbs.CreateAccount()
	require.NoError(t, err)

	// create client
	endpoint := getEndpoint()
	client := orbs.NewClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	contractName := deployContract(t, client, sender)

	// create transfer transaction
	tx, txId, err := client.CreateTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		contractName,
		"inc")
	require.NoError(t, err)

	// send the transaction
	incResponse, err := client.SendTransaction(tx)
	t.Log("Transfer response:")
	t.Logf("%+v", incResponse)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, incResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, incResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, incResponse.TransactionStatus)


	found, err := client.GetTransactionStatus(txId)
	require.NoError(t, err)
	require.EqualValues(t, codec.TRANSACTION_STATUS_COMMITTED, found.TransactionStatus)

	notFound, err := client.GetTransactionStatus("0xC0058950d1Bdde15d06C2d7354C3Cb15Dae02CFC6BF5934b358D43dEf1DFE1a0C420Da72e541bd6e")
	require.NoError(t, err)
	require.EqualValues(t, codec.TRANSACTION_STATUS_NO_RECORD_FOUND, notFound.TransactionStatus)
}
