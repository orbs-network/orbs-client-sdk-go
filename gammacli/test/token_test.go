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
	require.NoError(t, err, "get balance should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this

	out, err = runGammaCli("send-tx", "-i", "transfer.json")
	require.NoError(t, err, "transfer should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this

	txId := extractTxIdFromSendTxOutput(out)
	out, err = runGammaCli("status", "-txid", txId)
	require.NoError(t, err, "get tx status should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this

	out, err = runGammaCli("read", "-i", "get-balance.json")
	require.NoError(t, err, "get balance should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this
}
