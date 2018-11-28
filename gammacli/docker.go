package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const DOCKER_REPO = "orbsnetwork/gamma"
const DOCKER_RUN = "orbsnetwork/gamma"
const CONTAINER_NAME = "orbs-gamma-server"

func commandStartLocal() {
	gammaVersion := verifyDockerInstalled()

	if !doesFileExist(*flagKeyFile) {
		commandGenerateTestKeys()
	}

	if isDockerGammaRunning() {
		log(`
*********************************************************************************
              Orbs Gamma personal blockchain is already running!

  Run 'gamma-cli help' in terminal to learn how to interact with this instance.
              
**********************************************************************************
`)
		exit()
	}

	p := fmt.Sprintf("%d:8080", *flagPort)
	out, err := exec.Command("docker", "run", "-d", "--name", CONTAINER_NAME, "-p", p, DOCKER_RUN).CombinedOutput()
	if err != nil {
		die("could not exec 'docker run' command\n\n%s", out)
	}

	if !isDockerGammaRunning() {
		die("could not run docker image")
	}

	if *flagWait {
		waitUntilDockerIsReadyAndListening(IS_READY_TOTAL_WAIT_TIMEOUT)
	}

	log(`
*********************************************************************************
                 Orbs Gamma %s personal blockchain is running!

  Local blockchain instance started and listening on port %d.
  Run 'gamma-cli help' in terminal to learn how to interact with this instance.
              
**********************************************************************************
`, gammaVersion, *flagPort)
}

func commandStopLocal() {
	verifyDockerInstalled()

	if !isDockerGammaRunning() {
		die("gamma server instance is not started")
	}

	out, err := exec.Command("docker", "stop", CONTAINER_NAME).CombinedOutput()
	if err != nil {
		die("could not stop docker container\n\n%s", out)
	}

	out, err = exec.Command("docker", "rm", "-f", CONTAINER_NAME).CombinedOutput()
	if err != nil {
		die("could not rm docker container\n\n%s", out)
	}

	if isDockerGammaRunning() {
		die("could not stop docker container")
	}

	log(`
*********************************************************************************
                    Orbs Gamma personal blockchain stopped.

  A local blockchain instance is running in-memory.
  The next time you start the instance, all contracts and state will disappear. 
              
**********************************************************************************
`)
}

func commandUpgrade() {
	verifyDockerInstalled()

	log("downloading latest:\n")
	cmd := exec.Command("docker", "pull", DOCKER_REPO)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	log("")
}

func verifyDockerInstalled() string {
	out, err := exec.Command("docker", "images", DOCKER_REPO).CombinedOutput()
	if err != nil {
		if runtime.GOOS == "darwin" {
			die("docker is required but not installed on your machine\n\nInstall from:  https://docs.docker.com/docker-for-mac/install/")
		} else {
			die("docker is required but not installed on your machine\n\nInstall from:  https://docs.docker.com/install/")
		}
	}

	if strings.Count(string(out), "\n") > 1 {
		return extractTagFromDockerImagesOutput(string(out))
	}

	log("docker image not found, downloading:\n")
	cmd := exec.Command("docker", "pull", DOCKER_REPO)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	log("")

	out, err = exec.Command("docker", "images", DOCKER_REPO).CombinedOutput()
	if err != nil || strings.Count(string(out), "\n") == 1 {
		die("could not download docker image")
	}
	return extractTagFromDockerImagesOutput(string(out))
}

func isDockerGammaRunning() bool {
	out, err := exec.Command("docker", "ps", "-f", fmt.Sprintf("name=%s", CONTAINER_NAME)).CombinedOutput()
	if err != nil {
		die("could not exec 'docker ps' command\n\n%s", out)
	}
	return strings.Count(string(out), "\n") > 1
}

func extractTagFromDockerImagesOutput(out string) string {
	pattern := fmt.Sprintf(`%s\s+(\S+)`, regexp.QuoteMeta(DOCKER_REPO))
	re := regexp.MustCompile(pattern)
	res := re.FindStringSubmatch(out)
	if len(res) < 2 {
		return "unknown"
	}
	return res[1]
}
