package test

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
)

var cachedGammaCliBinaryPath string

func compileGammaCli() string {
	if cachedGammaCliBinaryPath != "" {
		return cachedGammaCliBinaryPath // cache compilation once per process
	}

	tempDir, err := ioutil.TempDir("", "gamma")
	if err != nil {
		panic(err)
	}

	binaryOutputPath := tempDir + "/gamma-cli"
	out, err := exec.Command("go", "build", "-o", binaryOutputPath, "../*.go").CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("compilation failed: %s\noutput:\n%s\n", err.Error(), out))
	}

	downloadLatestGammaServer(binaryOutputPath)

	cachedGammaCliBinaryPath = binaryOutputPath
	return cachedGammaCliBinaryPath
}

func downloadLatestGammaServer(gammaCliBinaryPath string) {
	out, err := exec.Command(gammaCliBinaryPath, "upgrade").CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("download latest gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	}
}

func runGammaCli(args ...string) (string, error) {
	out, err := exec.Command(compileGammaCli(), args...).CombinedOutput()
	return string(out), err
}

func startGammaServer() {
	out, err := runGammaCli("start-local")
	if err != nil {
		panic(fmt.Sprintf("start gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	}
}

func stopGammaServer() {
	runGammaCli("stop-local")
}

func extractTxIdFromSendTxOutput(out string) string {
	re := regexp.MustCompile(`\"TxId\":\s+\"(\w+)\"`)
	res := re.FindStringSubmatch(out)
	return res[1]
}
