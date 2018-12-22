package e2e

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/gammacli/test"
	"os"
)

const (
	GAMMA_PORT       = 8080
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


func getEndpoint() string {
	if endpoint := os.Getenv("GAMMA_ENDPOINT"); endpoint != "" {
		return endpoint
	}

	return fmt.Sprintf("http://%s:%d", GAMMA_ENDPOINT, GAMMA_PORT)
}