package main

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/gammacli/jsoncodec"
	"github.com/orbs-network/orbs-client-sdk-go/orbsclient"
	"io/ioutil"
	"strings"
)

const DEPLOY_SYSTEM_CONTRACT_NAME = "_Deployments"
const DEPLOY_SYSTEM_METHOD_NAME = "deployService"
const PROCESSOR_TYPE_NATIVE = uint32(1)
const PROCESSOR_TYPE_JAVASCRIPT = uint32(2)

func commandDeploy() {
	if *flagContractName == "" {
		die("contract name not provided, use the -name flag to provide it")
	}

	processorType := getProcessorTypeFromFilename(*flagCodeFile)
	code, err := ioutil.ReadFile(*flagCodeFile)
	if err != nil {
		die("could not open code file\n\n%s", err.Error())
	}

	signer := getTestKeyFromFile(*flagSigner)

	client := createOrbsClient()
	payload, txId, err := client.CreateSendTransactionPayload(signer.PublicKey, signer.PrivateKey, DEPLOY_SYSTEM_CONTRACT_NAME, DEPLOY_SYSTEM_METHOD_NAME, string(*flagContractName), uint32(processorType), []byte(code))
	if err != nil {
		die("could not encode payload of the message about to be sent to gamma server\n\n%s", err.Error())
	}
	response, err := client.SendTransaction(payload)
	if err != nil {
		die("failed sending transaction request to gamma server\n\n%s", err.Error())
	}

	output, err := jsoncodec.MarshalSendTxResponse(response, txId)
	if err != nil {
		die("could not encode response to json\n\n%s", err.Error())
	}

	log("%s\n", string(output))
}

func commandSendTx() {
	signer := getTestKeyFromFile(*flagSigner)

	bytes, err := ioutil.ReadFile(*flagInputFile)
	if err != nil {
		die("could not open input file\n\n%s", err.Error())
	}

	sendTx, err := jsoncodec.UnmarshalSendTx(bytes)
	if err != nil {
		die("failed parsing input json file '%s'\n\n%s", *flagInputFile, err.Error())
	}

	inputArgs, err := jsoncodec.UnmarshalArgs(sendTx.Arguments, getTestKeyFromFile)
	if err != nil {
		die(err.Error())
	}

	client := createOrbsClient()
	payload, txId, err := client.CreateSendTransactionPayload(signer.PublicKey, signer.PrivateKey, sendTx.ContractName, sendTx.MethodName, inputArgs...)
	if err != nil {
		die("could not encode payload of the message about to be sent to gamma server\n\n%s", err.Error())
	}

	response, clientErr := client.SendTransaction(payload)
	if response != nil {
		output, err := jsoncodec.MarshalSendTxResponse(response, txId)
		if err != nil {
			die("could not encode send-tx response to json\n\n%s", err.Error())
		}

		log("%s\n", string(output))
	}

	if clientErr != nil {
		die("send-tx request failed on gamma server\n\n%s", clientErr.Error())
	}
}

func commandRead() {
	signer := getTestKeyFromFile(*flagSigner)

	bytes, err := ioutil.ReadFile(*flagInputFile)
	if err != nil {
		die("could not open input file\n\n%s", err.Error())
	}

	read, err := jsoncodec.UnmarshalRead(bytes)
	if err != nil {
		die("failed parsing input json file '%s'\n\n%s", *flagInputFile, err.Error())
	}

	inputArgs, err := jsoncodec.UnmarshalArgs(read.Arguments, getTestKeyFromFile)
	if err != nil {
		die(err.Error())
	}

	client := createOrbsClient()
	payload, err := client.CreateCallMethodPayload(signer.PublicKey, read.ContractName, read.MethodName, inputArgs...)
	if err != nil {
		die("could not encode payload of the message about to be sent to gamma server\n\n%s", err.Error())
	}

	response, clientErr := client.CallMethod(payload)
	if response != nil {
		output, err := jsoncodec.MarshalReadResponse(response)
		if err != nil {
			die("could not encode read response to json\n\n%s", err.Error())
		}

		log("%s\n", string(output))
	}

	if clientErr != nil {
		die("read request failed on gamma server\n\n%s", clientErr.Error())
	}
}

func commandTxStatus() {
	if *flagTxId == "" {
		die("TxId not provided, it's given in the response of send-tx, use the -txid flag to provide it")
	}

	client := createOrbsClient()
	payload, err := client.CreateGetTransactionStatusPayload(*flagTxId)
	if err != nil {
		die("could not encode payload of the message about to be sent to gamma server\n\n%s", err.Error())
	}

	response, clientErr := client.GetTransactionStatus(payload)
	if response != nil {
		output, err := jsoncodec.MarshalTxStatusResponse(response)
		if err != nil {
			die("could not encode status response to json\n\n%s", err.Error())
		}

		log("%s\n", string(output))
	}

	if clientErr != nil {
		die("status request failed on gamma server\n\n%s", clientErr.Error())
	}
}

func createOrbsClient() *orbsclient.OrbsClient {
	env := getEnvironmentFromConfigFile(*flagEnv)
	if len(env.Endpoints) == 0 {
		die("environment Endpoints key does not contain any endpoints")
	}

	endpoint := env.Endpoints[0]
	if endpoint == "localhost" {
		if !isDockerGammaRunning() {
			die("local gamma server is not running, use gamma-cli start-local to start it")
		}
		endpoint = fmt.Sprintf("http://localhost:%d", *flagPort)
	}

	return orbsclient.NewOrbsClient(endpoint, env.VirtualChain, codec.NETWORK_TYPE_TEST_NET)
}

func getProcessorTypeFromFilename(filename string) uint32 {
	if strings.HasSuffix(filename, ".go") {
		return PROCESSOR_TYPE_NATIVE
	}
	if strings.HasSuffix(filename, ".js") {
		return PROCESSOR_TYPE_JAVASCRIPT
	}
	die("Unsupported code file type\n\nSupported code file extensions are: .go .js")
	return 0
}
