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
