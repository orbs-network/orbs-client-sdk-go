package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSimpleTransfer(t *testing.T) {
	startGammaServer()
	defer stopGammaServer()

	out, err := runGammaCli("read", "-i", "get-balance.json")
	t.Log(out)
	require.Error(t, err, "get balance should fail (not deployed)")
	require.True(t, strings.Contains(out, `"ExecutionResult": "ERROR_UNEXPECTED"`))

	out, err = runGammaCli("send-tx", "-i", "transfer.json")
	t.Log(out)
	require.NoError(t, err, "transfer should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	txId := extractTxIdFromSendTxOutput(out)
	t.Log(txId)

	// TODO: not sure why this isn't working, need to debug (difficult integration)
	//out, err = runGammaCli("status", "-txid", txId)
	//t.Log(out)
	//require.NoError(t, err, "get tx status should succeed")
	//require.True(t, strings.Contains(out, `"RequestStatus": "COMPLETED"`))

	out, err = runGammaCli("read", "-i", "get-balance.json")
	t.Log(out)
	require.NoError(t, err, "get balance should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "17"`))
}
