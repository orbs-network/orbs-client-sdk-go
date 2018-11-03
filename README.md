# Orbs Client SDK Go

> Client SDK for the Orbs blockchain in golang

## Usage

1. Create a few end user accounts:

    ```go
    sender, err := orbsclient.CreateAccount()
    receiver, err := orbsclient.CreateAccount()
    ```
    
2. Create a client instance:

    ```go
    const virtualChainId = 42
    client := orbsclient.NewOrbsClient("http://node-endpoint.com", virtualChainId, codec.NETWORK_TYPE_TEST_NET)
    ```

3. Send a transaction:

    ```go
    payload, txId, err := client.CreateSendTransactionPayload(
        sender.PublicKey,
        sender.PrivateKey,
        "BenchmarkToken",
        "transfer",
        uint64(10), receiver.RawAddress)
    response, err := client.SendTransaction(payload)
    ```
    
4. Check the transaction status:

    ```go
    payload, err = client.CreateGetTransactionStatusPayload(txId)
    response, err := client.GetTransactionStatus(payload)
    ```
    
5. Call a smart contract method:

    ```go
    payload, err = client.CreateCallMethodPayload(
        receiver.PublicKey,
        "BenchmarkToken",
        "getBalance",
        receiver.RawAddress)
    response, err := client.CallMethod(payload)
    ```

## Installation

1. Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  
    > Verify with `go version`

2. Get the library into your Go workspace:
 
     ```sh
     go get -u github.com/orbs-network/orbs-client-sdk-go/...
     ```

3. Import the client in your project: 

    ```go
    import "github.com/orbs-network/orbs-client-sdk-go/orbsclient" 
    ```

## Test

1. Run end to end test together with Gamma server:

    ```sh
    go test ./test/e2e
    ```

2. Run all tests:

    ```sh
    go test ./...
    ```

3. Run codec contract test (generate `output.json` as this is the reference implementation):

    ```sh
    cd ./test/codec/
    go run main.go
    ``` 

## Alternative Client SDK Implementations

Creating alternative client SDK implementations is highly encouraged, particularly in languages not supported by the core team. To make sure your client implementation is compliant to the Orbs [protocol specificiations](https://github.com/orbs-network/orbs-spec), we've created a set of compliance tests.

#### Codec contract test

The codec is the part that encodes and decodes protocol messages. The contract is created by the reference implementation (this repo) and embodied into an [input JSON](https://github.com/orbs-network/orbs-client-sdk-go/blob/master/test/codec/input.json) file and an [output JSON](https://github.com/orbs-network/orbs-client-sdk-go/blob/master/test/codec/output.json) file.

To run the contract test in your own implementation, encode/decode the messages according to `input.json` and compare your results to `output.json`. You can generate the JSON files from code by running

    ```sh
    cd ./test/codec/
    go run main.go
    ``` 

#### End to end test with Gamma server

Gamma server is a local development server for the Orbs network that can run on your own machine and be used for testing. This server can process smart contract deployments and transactions. The server is accessed via HTTP (just like a regular node) which makes it excellent for testing clients.

Take a look at the example [e2e test](https://github.com/orbs-network/orbs-client-sdk-go/tree/master/test/e2e) implemented in this repo.

#### Example compliant implementation

The JavaScript alternative implementation is a great source for inspiration. It contains both a codec [contract test](https://github.com/orbs-network/orbs-client-sdk-javascript/blob/master/src/codec/contract.test.ts) and an [end-to-end test](https://github.com/orbs-network/orbs-client-sdk-javascript/tree/master/e2e/nodejs).
