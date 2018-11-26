package main

import (
	"flag"
	"fmt"
	"os"
)

const GAMMA_CLI_VERSION = "0.5.0"
const CONFIG_FILENAME = "orbs-gamma-config.json"
const TEST_KEYS_FILENAME = "orbs-test-keys.json"
const LOCAL_ENV_ID = "local"

type command struct {
	desc    string
	args    string
	example string
	handler func()
	sort    int
}

var commands = map[string]*command{
	"start-local": {
		desc:    "start a local Orbs personal blockchain instance listening on port",
		args:    "-port <PORT> -wait",
		example: "gamma-cli start-local -port 8080",
		handler: commandStartLocal,
		sort:    0,
	},
	"stop-local": {
		desc:    "stop a locally running Orbs personal blockchain instance",
		handler: commandStopLocal,
		sort:    1,
	},
	"gen-test-keys": {
		desc:    "generate a new batch of 10 test keys and store in " + TEST_KEYS_FILENAME + " (default filename)",
		args:    "-keys <OUTPUT_FILE>",
		example: "gamma-cli gen-test-keys -keys " + TEST_KEYS_FILENAME,
		handler: commandGenerateTestKeys,
		sort:    2,
	},
	"deploy": {
		desc:    "deploy a smart contract with the code specified in contract.go (default filename)",
		args:    "-name <CONTRACT_NAME> -code <CODE_FILE> -signer <ID_FROM_KEYS_JSON>",
		example: "gamma-cli deploy -name MyToken -code contract.go -signer user1",
		handler: commandDeploy,
		sort:    3,
	},
	"send-tx": {
		desc:    "sign and send the transaction specified in input.json (default filename)",
		args:    "-i <INPUT_FILE> -signer <ID_FROM_KEYS_JSON>",
		example: "gamma-cli send-tx -i transfer.json -signer user1",
		handler: commandSendTx,
		sort:    4,
	},
	"read": {
		desc:    "read state or run a read-only contract method as specified in input.json (default filename)",
		args:    "-i <INPUT_FILE> -signer <ID_FROM_KEYS_JSON>",
		example: "gamma-cli read -i get-balance.json -signer user1",
		handler: commandRead,
		sort:    5,
	},
	"status": {
		desc:    "get the current status of a sent transaction",
		args:    "-txid <TX_ID>",
		example: "gamma-cli status -txid nXAmGL2peGvXkrDxC2cFaZwhykfMGFGj1DUJ9eDFRdSnNgCpQ69MQz",
		handler: commandTxStatus,
		sort:    6,
	},
	"upgrade": {
		desc:    "upgrade to the latest version of gamma server",
		handler: commandUpgrade,
		sort:    7,
	},
	"version": {
		desc:    "print gamma-cli and gamma server versions",
		handler: commandVersion,
		sort:    8,
	},
	"help": {
		desc: "print this help screen",
		sort: 9,
	},
}

var (
	flagPort         = flag.Int("port", 8080, "listening port for gamma server")
	flagSigner       = flag.String("signer", "user1", "id of the signing key from the test key json")
	flagContractName = flag.String("name", "", "name of the smart contract being deployed")
	flagCodeFile     = flag.String("code", "contract.go", "source file for the smart contract being deployed, normally .go or .js file")
	flagInputFile    = flag.String("i", "input.json", "name of the json input file")
	flagKeyFile      = flag.String("keys", TEST_KEYS_FILENAME, "name of the json file containing test keys")
	flagTxId         = flag.String("txid", "", "TxId of a previously sent transaction, given in the response of send-tx")
	flagEnv          = flag.String("env", LOCAL_ENV_ID, fmt.Sprintf("environment from %s containing server connection details", CONFIG_FILENAME))
	flagWait         = flag.Bool("wait", false, "wait until gamma server is ready and listening")
)

func main() {
	flag.Usage = commandShowHelp
	commands["help"].handler = commandShowHelp

	if len(os.Args) <= 1 {
		commandShowHelp()
	}
	cmdName := os.Args[1]
	os.Args = os.Args[1:]
	flag.Parse()

	cmd, found := commands[cmdName]
	if !found {
		die("command '%s' not found, run 'gamma-cli help' to see available commands", cmdName)
	}

	cmd.handler()
}

func log(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintf(os.Stdout, "\n")
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR:\n  ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n\n")
	os.Exit(2)
}

func exit() {
	os.Exit(0)
}

func doesFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
