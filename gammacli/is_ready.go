package main

import (
	"fmt"
	"net"
	"time"
)

const IS_READY_TOTAL_WAIT_TIMEOUT = 20 * time.Second
const IS_READY_POLLING_INTERVAL = 500 * time.Millisecond

const DEPLOY_GET_INFO_SYSTEM_METHOD_NAME = "getInfo"

func isDockerReadyAndListening() bool {
	signer := getTestKeyFromFile(*flagSigner)

	client := createOrbsClient()
	payload, err := client.CreateCallMethodPayload(signer.PublicKey, DEPLOY_SYSTEM_CONTRACT_NAME, DEPLOY_GET_INFO_SYSTEM_METHOD_NAME, DEPLOY_SYSTEM_CONTRACT_NAME)
	if err != nil {
		die("Could not encode payload of the message about to be sent to Gamma server.\n\n%s", err.Error())
	}

	_, err = client.CallMethod(payload)
	return err == nil
}

func waitUntilDockerIsReadyAndListening(timeout time.Duration) {
	start := time.Now()
	for time.Now().Sub(start) < timeout {
		if isDockerReadyAndListening() {
			return
		}
		time.Sleep(IS_READY_POLLING_INTERVAL)
	}
}

func isPortListening(port int) bool {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true // if it fails then the port is likely taken
	}
	server.Close()
	return false
}
