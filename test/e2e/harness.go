package e2e

import (
	"flag"
	"github.com/orbs-network/orbs-network-go/devtools/gammacli"
	"strconv"
)

type harness struct {
	gamma *gammacli.GammaServer
	port  int
}

func (h *harness) shutdown() {
	h.gamma.GracefulShutdown(0) // meaning don't have a deadline timeout so allowing enough time for shutdown to free port
}

func newHarness(host string, port int) *harness {
	p := flag.Int("port", port, "The port to bind the gamma-server to")
	flag.Parse()

	var serverAddress = host + ":" + strconv.Itoa(*p)

	server := gammacli.StartGammaServer(serverAddress, false)
	return &harness{gamma: server, port: server.Port()}
}