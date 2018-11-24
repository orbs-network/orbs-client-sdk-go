package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const DOCKER_REPO = "orbs"
const DOCKER_RUN = "orbs:gamma-server"
const CONTAINER_NAME = "orbs-gamma-server"

func commandStartLocal() {
	verifyDockerInstalled()

	if isDockerGammaRunning() {
		log(`
*********************************************************************************
                 Orbs Gamma personal blockchain is running!

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

	log(`
*********************************************************************************
                 Orbs Gamma v%s personal blockchain is running!

  Local blockchain instance started and listening on port %d.
  Run 'gamma-cli help' in terminal to learn how to interact with this instance.
              
**********************************************************************************
`, GAMMA_CLI_VERSION, *flagPort) // TODO: change version to gamma version (from the docker tag)
}

func commandStopLocal() {
	verifyDockerInstalled()

	if !isDockerGammaRunning() {
		die("gamma server instance is not started")
	}

	out, err := exec.Command("docker", "rm", "-f", CONTAINER_NAME).CombinedOutput()
	if err != nil {
		die("could not stop docker container\n\n%s", out)
	}

	if isDockerGammaRunning() {
		die("could not stop docker container")
	}

	log(`
*********************************************************************************
                    Orbs Gamma personal blockchain stopped.

  The local blockchain instance was only running in memory.
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

func verifyDockerInstalled() {
	out, err := exec.Command("docker", "images", DOCKER_REPO).CombinedOutput()
	if err != nil {
		if runtime.GOOS == "darwin" {
			die("docker is required but not installed on your machine\n\nInstall from:  https://docs.docker.com/docker-for-mac/install/")
		} else {
			die("docker is required but not installed on your machine\n\nInstall from:  https://docs.docker.com/install/")
		}
	}

	if strings.Count(string(out), "\n") > 1 {
		return
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
}

func isDockerGammaRunning() bool {
	out, err := exec.Command("docker", "ps", "-f", fmt.Sprintf("name=%s", CONTAINER_NAME)).CombinedOutput()
	if err != nil {
		die("could not exec 'docker ps' command\n\n%s", out)
	}
	return strings.Count(string(out), "\n") > 1
}
