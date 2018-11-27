package test

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

var cachedGammaCliBinaryPath string

type gammaCli struct {
	port string
}

func compileGammaCli() string {
	if cachedGammaCliBinaryPath != "" {
		return cachedGammaCliBinaryPath // cache compilation once per process
	}

	tempDir, err := ioutil.TempDir("", "gamma")
	if err != nil {
		panic(err)
	}

	binaryOutputPath := tempDir + "/gamma-cli"
	sourcePackage := "github.com/orbs-network/orbs-client-sdk-go/gammacli"
	out, err := exec.Command("go", "build", "-o", binaryOutputPath, sourcePackage).CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("compilation failed: %s\noutput:\n%s\n", err.Error(), out))
	} else {
		fmt.Printf("compiled gamma-cli successfully:\n %s\n", binaryOutputPath)
	}

	downloadLatestGammaServer(binaryOutputPath)

	cachedGammaCliBinaryPath = binaryOutputPath
	return cachedGammaCliBinaryPath
}

func downloadLatestGammaServer(gammaCliBinaryPath string) {
	start := time.Now()
	// TODO: bring this part back and improve its success
	//out, err := exec.Command(gammaCliBinaryPath, "upgrade").CombinedOutput()
	//if err != nil {
	//	panic(fmt.Sprintf("download latest gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	//}
	delta := time.Now().Sub(start)
	fmt.Printf("upgraded gamma-server to latest version (this took %.3fs)\n", delta.Seconds())
}

func (g *gammaCli) Run(args ...string) (string, error) {
	if len(args) > 0 {
		args = append(args, "-port", g.port)
	}
	out, err := exec.Command(compileGammaCli(), args...).CombinedOutput()
	return string(out), err
}

func GammaCli() *gammaCli {
	return GammaCliWithPort(8080)
}

func GammaCliWithPort(port int) *gammaCli {
	return &gammaCli{
		port: fmt.Sprintf("%d", port),
	}
}

func (g *gammaCli) StartGammaServer() *gammaCli {
	out, err := g.Run("start-local", "-wait")
	if err != nil {
		panic(fmt.Sprintf("start gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	}
	return g
}

func (g *gammaCli) StopGammaServer() {
	g.Run("stop-local")
}

func extractTxIdFromSendTxOutput(out string) string {
	re := regexp.MustCompile(`\"TxId\":\s+\"(\w+)\"`)
	res := re.FindStringSubmatch(out)
	return res[1]
}

func getCurrentSourceFileDirPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
