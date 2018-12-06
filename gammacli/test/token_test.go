package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSimpleTransfer(t *testing.T) {
	cli := GammaCli().DownloadLatestGammaServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("read", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-i", "get-balance.json")
	t.Log(out)
	require.Error(t, err, "get balance should fail (not deployed)")
	require.True(t, strings.Contains(out, `"ExecutionResult": "ERROR_UNEXPECTED"`))

	out, err = cli.Run("send-tx", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-i", "transfer.json")
	t.Log(out)
	require.NoError(t, err, "transfer should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	txId := extractTxIdFromSendTxOutput(out)
	t.Log(txId)

	out, err = cli.Run("status", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-txid", txId)
	t.Log(out)
	require.NoError(t, err, "get tx status should succeed")
	require.True(t, strings.Contains(out, `"RequestStatus": "COMPLETED"`))

	out, err = cli.Run("read", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-i", "get-balance.json")
	t.Log(out)
	require.NoError(t, err, "get balance should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "17"`))
}
