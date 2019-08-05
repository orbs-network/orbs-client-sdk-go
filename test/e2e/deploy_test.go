package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-client-sdk-go/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeployMultifile(t *testing.T) {
	h := newHarness()
	defer h.shutdown()

	// create sender account
	sender, err := orbs.CreateAccount()
	require.NoError(t, err)

	// create client
	endpoint := getEndpoint()
	client := orbs.NewClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	sources, err := utils.ReadSourcesFromDir("./contract")
	require.NoError(t, err)
	require.Len(t, sources, 2)

	// create transfer transaction
	deployTx, _, err := client.CreateDeployTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		"Inc",
		orbs.PROCESSOR_TYPE_NATIVE,
		sources...)
	require.NoError(t, err)

	deployResponse, err := client.SendTransaction(deployTx)

	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, deployResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, deployResponse.TransactionStatus)

	// create transfer transaction
	tx, _, err := client.CreateTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		"Inc",
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
}