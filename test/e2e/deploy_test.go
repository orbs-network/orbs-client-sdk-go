package e2e

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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

	contractName := deployContract(t, client, sender)

	// create transfer transaction
	tx, _, err := client.CreateTransaction(
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
}

func deployContract(t *testing.T, client *orbs.OrbsClient, sender *orbs.OrbsAccount) string {
	sources, err := orbs.ReadSourcesFromDir("./contract")
	require.NoError(t, err)
	require.Len(t, sources, 2)

	contractName := fmt.Sprintf("Inc%d", time.Now().UnixNano())

	// create transfer transaction
	deployTx, _, err := client.CreateDeployTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		contractName,
		orbs.PROCESSOR_TYPE_NATIVE,
		sources...)
	require.NoError(t, err)
	deployResponse, err := client.SendTransaction(deployTx)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, deployResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, deployResponse.TransactionStatus)

	return contractName
}