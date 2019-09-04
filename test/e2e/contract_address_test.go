package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContractAddress(t *testing.T) {
	h := newHarness()
	defer h.shutdown()

	// create sender account
	sender, err := orbs.CreateAccount()
	require.NoError(t, err)

	// create client
	endpoint := getEndpoint()
	client := orbs.NewClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	contractName := deployContract(t, client, sender)

	query, err := client.CreateQuery(sender.PublicKey, contractName, "getOwnAddress")
	require.NoError(t, err)

	response, err := client.SendQuery(query)
	require.NoError(t, err)
	require.EqualValues(t, response.OutputArguments[0], orbs.ContractNameToAddressAsBytes(contractName))
}