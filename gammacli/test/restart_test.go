package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRestart(t *testing.T) {
	cli := GammaCli()
	defer cli.StopGammaServer()

	_, err := cli.Run("start-local")
	require.NoError(t, err, "start Gamma server should succeed")

	_, err = cli.Run("stop-local")
	require.NoError(t, err, "stop Gamma server should succeed")

	_, err = cli.Run("stop-local")
	require.NoError(t, err, "second stop Gamma server should succeed")

	_, err = cli.Run("start-local")
	require.NoError(t, err, "start Gamma server should succeed")
}
