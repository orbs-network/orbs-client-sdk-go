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
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this

	out, err = runGammaCli("read", "-i", "counter-get.json")
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this

	out, err = runGammaCli("send-tx", "-i", "counter-add.json")
	require.NoError(t, err, "add should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this

	out, err = runGammaCli("read", "-i", "counter-get.json")
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, "TODO")) // TODO: finish this
}
