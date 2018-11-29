package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

const DOCKER_REPO = "orbsnetwork/gamma"
const DOCKER_RUN = "orbsnetwork/gamma:%s"
const CONTAINER_NAME = "orbs-gamma-server"
const DOCKER_REGISTRY_TAGS_URL = "https://registry.hub.docker.com/v2/repositories/orbsnetwork/gamma/tags/"

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
	run := fmt.Sprintf(DOCKER_RUN, gammaVersion)
	out, err := exec.Command("docker", "run", "-d", "--name", CONTAINER_NAME, "-p", p, run).CombinedOutput()
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

func commandUpgradeServer() {
	currentTag := verifyDockerInstalled()
	latestTag := getLatestDockerTag()

	if cmpTags(latestTag, currentTag) <= 0 {
		log("current version %s does not require upgrade\n", currentTag)
		exit()
	}

	log("downloading latest version %s:\n", latestTag)
	cmd := exec.Command("docker", "pull", fmt.Sprintf("%s:%s", DOCKER_REPO, latestTag))
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
		return false
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

func getLatestDockerTag() string {
	resp, err := http.Get(DOCKER_REGISTRY_TAGS_URL)
	if err != nil {
		die("cannot connect to docker registry to get image list")
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(bytes) == 0 {
		die("bad image list response from docker registry")
	}
	tag, err := extractLatestTagFromDockerHubResponse(bytes)
	if err != nil {
		die("cannot parse image list response from docker registry")
	}
	return tag
}

type dockerHubTagsJson struct {
	Results []*struct {
		Name string
	}
}

func extractLatestTagFromDockerHubResponse(responseBytes []byte) (string, error) {
	var response *dockerHubTagsJson
	err := json.Unmarshal(responseBytes, &response)
	if err != nil {
		return "", err
	}
	maxTag := ""
	for _, result := range response.Results {
		if cmpTags(result.Name, maxTag) > 0 {
			maxTag = result.Name
		}
	}
	if maxTag == "" {
		return "", errors.New("no valid tags found")
	}
	return maxTag, nil
}

func cmpTags(t1, t2 string) int {
	re := regexp.MustCompile(`v(\d+)\.(\d+)\.(\d+)`)
	m1 := re.FindStringSubmatch(t1)
	if len(m1) < 4 {
		return -1
	}
	m2 := re.FindStringSubmatch(t2)
	if len(m2) < 4 {
		return 1
	}
	diff := atoi(m1[1]) - atoi(m2[1])
	if diff != 0 {
		return diff
	}
	diff = atoi(m1[2]) - atoi(m2[2])
	if diff != 0 {
		return diff
	}
	return atoi(m1[3]) - atoi(m2[3])
}

func atoi(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return res
}