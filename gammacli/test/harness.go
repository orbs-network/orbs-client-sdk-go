package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

const CONFIG_FILENAME = "orbs-gamma-config.json"

var cachedGammaCliBinaryPath string
var downloadedLatestGammaServer bool

type gammaCli struct {
	port string
}

func compileGammaCli() string {
	if gammaCliPathFromEnv := os.Getenv("GAMMA_CLI_PATH"); gammaCliPathFromEnv != "" {
		return gammaCliPathFromEnv
	}

	if cachedGammaCliBinaryPath != "" {
		return cachedGammaCliBinaryPath // cache compilation once per process
	}

	tempDir, err := ioutil.TempDir("", "gamma")
	if err != nil {
		panic(err)
	}

	binaryOutputPath := tempDir + "/gamma-cli"
	cmd := exec.Command("go", "build", "-o", binaryOutputPath, ".")
	cmd.Dir = path.Join(getCurrentSourceFileDirPath(), "..")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("compilation failed: %s\noutput:\n%s\n", err.Error(), out))
	} else {
		fmt.Printf("compiled gamma-cli successfully:\n %s\n", binaryOutputPath)
	}

	cachedGammaCliBinaryPath = binaryOutputPath
	return cachedGammaCliBinaryPath
}

func (g *gammaCli) Run(args ...string) (string, error) {
	if len(args) > 0 {
		// streamlined supprt for a different port
		args = append(args, "-port", g.port)

		// needed for dockerized tests in ci
		if args[0] != "start-local" && args[0] != "stop-local" && args[0] != "upgrade-server" {
			args = append(args, "-env", getGammaEnvironmentFromEnvVar())
			pathToConfig := path.Join(os.Getenv("HOME"), ".orbs", CONFIG_FILENAME)
			args = append(args, "-config", pathToConfig)
		}
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
		panic(fmt.Sprintf("start Gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	}
	fmt.Println(out)
	return g
}

func (g *gammaCli) StopGammaServer() {
	out, err := g.Run("stop-local")
	if err != nil {
		panic(fmt.Sprintf("stop Gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	}
}

func (g *gammaCli) DownloadLatestGammaServer() *gammaCli {
	if downloadedLatestGammaServer {
		return g
	}
	downloadedLatestGammaServer = true

	start := time.Now()
	out, err := g.Run("upgrade-server")
	if err != nil {
		panic(fmt.Sprintf("download latest Gamma server failed: %s\noutput:\n%s\n", err.Error(), out))
	}
	delta := time.Now().Sub(start)
	fmt.Printf("upgraded gamma-server to latest version (this took %.3fs)\n", delta.Seconds())
	return g
}

func getGammaEnvironmentFromEnvVar() string {
	if env := os.Getenv("GAMMA_ENVIRONMENT"); env != "" {
		return env
	}

	return "local"
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
