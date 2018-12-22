package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDeployCounter(t *testing.T) {
	cli := GammaCli().DownloadLatestGammaServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("deploy", "-name", "CounterExample", "-code", "./_counter/contract.go")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = cli.Run("read", "-i", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "0"`))

	out, err = cli.Run("send-tx", "-i", "counter-add.json")
	t.Log(out)
	require.NoError(t, err, "add should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "previous count is 0"`))

	out, err = cli.Run("read", "-i", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "25"`))
}
