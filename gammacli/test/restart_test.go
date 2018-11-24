package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRestart(t *testing.T) {
	defer stopGammaServer()

	_, err := runGammaCli("start-local")
	require.NoError(t, err, "start gamma server should succeed")

	_, err = runGammaCli("stop-local")
	require.NoError(t, err, "stop gamma server should succeed")

	_, err = runGammaCli("stop-local")
	require.Error(t, err, "second stop gamma server should fail")

	_, err = runGammaCli("start-local")
	require.NoError(t, err, "start gamma server should succeed")
}
