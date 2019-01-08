package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbsclient"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestTextualError(t *testing.T) {
	h := newHarness()
	defer h.shutdown()

	// create client
	endpoint := getEndpoint()
	client := orbsclient.NewOrbsClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	// send a corrupt transaction
	transferResponse, err := client.SendTransaction([]byte{0x01, 0x02, 0x03})
	require.Error(t, err, "request should fail")
	require.True(t, strings.Contains(err.Error(), "http request is not a valid membuffer"))
	require.Nil(t, transferResponse, "response should be nil")
}
