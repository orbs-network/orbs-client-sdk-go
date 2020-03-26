package e2e

import (
	"context"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
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
	tx, _, err := client.CreateTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		contractName,
		"inc")
	require.NoError(t, err)

	var events []*codec.Event
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go client.Subscribe(ctx, sender.PublicKey, contractName, []string{"Inc"}, 5*time.Millisecond, func(event *codec.Event, blockHeight uint64, txHash []byte, txId []byte, eventIndex uint64) error {
		events = append(events, event)
		return nil
	})

	// send the transaction
	incResponse, err := client.SendTransaction(tx)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, incResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, incResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, incResponse.TransactionStatus)

	time.Sleep(1*time.Second)

	require.EqualValues(t, &codec.Event{
		ContractName: contractName,
		EventName:    "Inc",
		Arguments: []interface{}{
			uint64(1),
		},
	}, events[0])
}
