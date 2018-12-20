package main

import (
	"encoding/json"
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

// the purpose of the contract test is to create an output.json which other client implementations can compare against
// this is the reference client implementations

type ScenarioInput struct {
	Test                               string
	SendTransactionRequest             *codec.SendTransactionRequest
	CallMethodRequest                  *codec.CallMethodRequest
	GetTransactionStatusRequest        *codec.GetTransactionStatusRequest
	GetTransactionReceiptProofRequest  *codec.GetTransactionReceiptProofRequest
	PrivateKey                         []byte
	InputArgumentsTypes                []string
	SendTransactionResponse            []byte
	CallMethodResponse                 []byte
	GetTransactionStatusResponse       []byte
	GetTransactionReceiptProofResponse []byte
}

type ScenarioOutput struct {
	Test                               string
	SendTransactionRequest             []byte                                    `json:",omitempty"`
	CallMethodRequest                  []byte                                    `json:",omitempty"`
	GetTransactionStatusRequest        []byte                                    `json:",omitempty"`
	GetTransactionReceiptProofRequest  []byte                                    `json:",omitempty"`
	TxId                               []byte                                    `json:",omitempty"`
	SendTransactionResponse            *codec.SendTransactionResponse            `json:",omitempty"`
	CallMethodResponse                 *codec.CallMethodResponse                 `json:",omitempty"`
	GetTransactionStatusResponse       *codec.GetTransactionStatusResponse       `json:",omitempty"`
	GetTransactionReceiptProofResponse *codec.GetTransactionReceiptProofResponse `json:",omitempty"`
}

func die(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
	os.Exit(-1)
}

func main() {
	// read input.json
	inputBytes, err := ioutil.ReadFile("input.json")
	if err != nil {
		die(err)
	}

	// unmarshal input
	var input []*ScenarioInput
	err = json.Unmarshal(inputBytes, &input)
	if err != nil {
		die(err)
	}

	// go over all scenarios
	var output []*ScenarioOutput
	for index, scenarioInput := range input {
		scenarioOutput, err := generateOutput(scenarioInput)
		if err != nil {
			die(errors.Wrapf(err, "scenario index %d failed to run\n%+v", index, scenarioInput))
		}
		output = append(output, scenarioOutput)
	}

	// marshal output
	outputBytes, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		die(err)
	}

	// write output.json
	err = ioutil.WriteFile("output.json", outputBytes, 0655)
	if err != nil {
		die(err)
	}

	println("output.json written successfully\n")
}

func generateOutput(scenarioInput *ScenarioInput) (*ScenarioOutput, error) {
	// SendTransactionRequest
	if scenarioInput.SendTransactionRequest != nil {
		scenarioInput.SendTransactionRequest.FixInputArgumentTypes(scenarioInput.InputArgumentsTypes)
		encodedBytes, txId, err := codec.EncodeSendTransactionRequest(scenarioInput.SendTransactionRequest, scenarioInput.PrivateKey)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, SendTransactionRequest: encodedBytes, TxId: txId}, nil
	}

	// CallMethodRequest
	if scenarioInput.CallMethodRequest != nil {
		scenarioInput.CallMethodRequest.FixInputArgumentTypes(scenarioInput.InputArgumentsTypes)
		encodedBytes, err := codec.EncodeCallMethodRequest(scenarioInput.CallMethodRequest)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, CallMethodRequest: encodedBytes}, nil
	}

	// GetTransactionStatusRequest
	if scenarioInput.GetTransactionStatusRequest != nil {
		encodedBytes, err := codec.EncodeGetTransactionStatusRequest(scenarioInput.GetTransactionStatusRequest)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, GetTransactionStatusRequest: encodedBytes}, nil
	}

	// GetTransactionReceiptProofRequest
	if scenarioInput.GetTransactionReceiptProofRequest != nil {
		encodedBytes, err := codec.EncodeGetTransactionReceiptProofRequest(scenarioInput.GetTransactionReceiptProofRequest)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, GetTransactionReceiptProofRequest: encodedBytes}, nil
	}

	// SendTransactionResponse
	if scenarioInput.SendTransactionResponse != nil {
		res, err := codec.DecodeSendTransactionResponse(scenarioInput.SendTransactionResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, SendTransactionResponse: res}, nil
	}

	// CallMethodResponse
	if scenarioInput.CallMethodResponse != nil {
		res, err := codec.DecodeCallMethodResponse(scenarioInput.CallMethodResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, CallMethodResponse: res}, nil
	}

	// GetTransactionStatusResponse
	if scenarioInput.GetTransactionStatusResponse != nil {
		res, err := codec.DecodeGetTransactionStatusResponse(scenarioInput.GetTransactionStatusResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, GetTransactionStatusResponse: res}, nil
	}

	// GetTransactionReceiptProofResponse
	if scenarioInput.GetTransactionReceiptProofResponse != nil {
		res, err := codec.DecodeGetTransactionReceiptProofResponse(scenarioInput.GetTransactionReceiptProofResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, GetTransactionReceiptProofResponse: res}, nil
	}

	return nil, errors.New("scenario type unrecognized")
}
