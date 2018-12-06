package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDeployCounter(t *testing.T) {
	cli := GammaCli().DownloadLatestGammaServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("deploy", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-name", "CounterExample", "-code", "./counter/contract.go")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = cli.Run("read", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-i", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "0"`))

	out, err = cli.Run("send-tx", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-i", "counter-add.json")
	t.Log(out)
	require.NoError(t, err, "add should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = cli.Run("read", "-env", getGammaEnvironment(), "-config", "./orbs-gamma-config.json", "-i", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "25"`))
}
