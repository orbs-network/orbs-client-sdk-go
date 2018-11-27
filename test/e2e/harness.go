package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/gammacli/test"
)

const (
	GAMMA_PORT       = 8092
	GAMMA_ENDPOINT   = "localhost"
	VIRTUAL_CHAIN_ID = 42 // gamma-cli config default
)

type harness struct {
}

func (h *harness) shutdown() {
	test.GammaCliWithPort(GAMMA_PORT).StopGammaServer()
}

func newHarness() *harness {
	test.GammaCliWithPort(GAMMA_PORT).StartGammaServer()
	return &harness{}
}
