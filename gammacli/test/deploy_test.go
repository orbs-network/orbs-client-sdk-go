package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDeployCounter(t *testing.T) {
	startGammaServer()
	defer stopGammaServer()

	out, err := runGammaCli("deploy", "-name", "CounterExample", "-code", "./counter/contract.go")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = runGammaCli("read", "-i", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "0"`))

	out, err = runGammaCli("send-tx", "-i", "counter-add.json")
	t.Log(out)
	require.NoError(t, err, "add should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = runGammaCli("read", "-i", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "25"`))
}
