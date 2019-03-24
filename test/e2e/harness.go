// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	GAMMA_PORT       = 8080
	GAMMA_ENDPOINT   = "localhost"
	VIRTUAL_CHAIN_ID = 42 // gamma-cli config default
	EXPERIMENTAL     = true
)

type harness struct {
}

func (h *harness) gammaCliRun(args []string) {
	args = append(args, "-port", strconv.Itoa(GAMMA_PORT))
	if EXPERIMENTAL {
		args = append(args, "-env", "experimental")
	}
	fmt.Printf("RUNNING: gamma-cli %s\n", strings.Join(args, " "))
	out, err := exec.Command("gamma-cli", args...).CombinedOutput()
	if len(out) > 0 {
		fmt.Printf("%s\n", string(out))
	}
	if err != nil {
		panic("Unable to run E2E, make sure gamma-cli is installed (https://github.com/orbs-network/gamma-cli). Error: " + err.Error())
	}
}

func newHarness() *harness {
	h := &harness{}
	h.gammaCliRun([]string{"start-local", "-wait"})
	return h
}

func (h *harness) shutdown() {
	h.gammaCliRun([]string{"stop-local"})
}

func getEndpoint() string {
	if endpoint := os.Getenv("GAMMA_ENDPOINT"); endpoint != "" {
		return endpoint
	}

	return fmt.Sprintf("http://%s:%d", GAMMA_ENDPOINT, GAMMA_PORT)
}
