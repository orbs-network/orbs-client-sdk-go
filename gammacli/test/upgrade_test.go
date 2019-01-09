package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestUpgradeServer(t *testing.T) {
	cli := GammaCli()

	out, err := cli.Run("version")
	t.Log(out)
	require.NoError(t, err, "version should succeed")
	require.True(t, strings.Contains(out, `(docker)`), "version output should show docker")

	out, err = cli.Run("upgrade-server")
	t.Log(out)
	require.NoError(t, err, "upgrade server experimental should succeed")
	require.True(t, strings.Contains(out, `does not require upgrade`), "upgrade same tag should not try to pull fresh copy")
}

func TestUpgradeExperimentalServer(t *testing.T) {
	cli := GammaCli()
	defer cli.StopGammaServer()

	out, err := cli.Run("version", "-experimental")
	t.Log(out)
	require.NoError(t, err, "version experimental should succeed")
	require.True(t, strings.Contains(out, `experimental (docker)`), "version output should show experimental docker")

	out, err = cli.Run("upgrade-server", "-experimental")
	t.Log(out)
	require.NoError(t, err, "upgrade server experimental should succeed")
	require.True(t, strings.Contains(out, `experimental: Pulling from orbsnetwork/gamma`), "experimental upgrade should always try to pull fresh copy")

	out, err = cli.Run("start-local", "-experimental")
	t.Log(out)
	require.NoError(t, err, "start Gamma server should succeed")
	require.True(t, strings.Contains(out, `Orbs Gamma experimental personal blockchain`), "started Gamma server should not be experimental")
}
