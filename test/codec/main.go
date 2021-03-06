// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

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
	RunQueryRequest                    *codec.RunQueryRequest
	GetTransactionStatusRequest        *codec.GetTransactionStatusRequest
	GetTransactionReceiptProofRequest  *codec.GetTransactionReceiptProofRequest
	GetBlockRequest                    *codec.GetBlockRequest
	PrivateKey                         []byte
	InputArgumentsTypes                []string
	SendTransactionResponse            []byte
	RunQueryResponse                   []byte
	GetTransactionStatusResponse       []byte
	GetTransactionReceiptProofResponse []byte
	GetBlockResponse                   []byte
}

type ScenarioOutput struct {
	Test                               string
	SendTransactionRequest             []byte                                    `json:",omitempty"`
	RunQueryRequest                    []byte                                    `json:",omitempty"`
	GetTransactionStatusRequest        []byte                                    `json:",omitempty"`
	GetTransactionReceiptProofRequest  []byte                                    `json:",omitempty"`
	GetBlockRequest                    []byte                                    `json:",omitempty"`
	TxId                               []byte                                    `json:",omitempty"`
	SendTransactionResponse            *codec.SendTransactionResponse            `json:",omitempty"`
	RunQueryResponse                   *codec.RunQueryResponse                   `json:",omitempty"`
	GetTransactionStatusResponse       *codec.GetTransactionStatusResponse       `json:",omitempty"`
	GetTransactionReceiptProofResponse *codec.GetTransactionReceiptProofResponse `json:",omitempty"`
	GetBlockResponse                   *codec.GetBlockResponse                   `json:",omitempty"`
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
		encodedBytes, txId, err := codec.EncodeSendTransactionRequest(scenarioInput.SendTransactionRequest, scenarioInput.PrivateKey)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, SendTransactionRequest: encodedBytes, TxId: txId}, nil
	}

	// RunQueryRequest
	if scenarioInput.RunQueryRequest != nil {
		encodedBytes, err := codec.EncodeRunQueryRequest(scenarioInput.RunQueryRequest)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, RunQueryRequest: encodedBytes}, nil
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

	// GetBlockRequest
	if scenarioInput.GetBlockRequest != nil {
		encodedBytes, err := codec.EncodeGetBlockRequest(scenarioInput.GetBlockRequest)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, GetBlockRequest: encodedBytes}, nil
	}

	// SendTransactionResponse
	if scenarioInput.SendTransactionResponse != nil {
		res, err := codec.DecodeSendTransactionResponse(scenarioInput.SendTransactionResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, SendTransactionResponse: res}, nil
	}

	// RunQueryResponse
	if scenarioInput.RunQueryResponse != nil {
		res, err := codec.DecodeRunQueryResponse(scenarioInput.RunQueryResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, RunQueryResponse: res}, nil
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

	// GetBlockResponse
	if scenarioInput.GetBlockResponse != nil {
		res, err := codec.DecodeGetBlockResponse(scenarioInput.GetBlockResponse)
		if err != nil {
			return nil, err
		}
		return &ScenarioOutput{Test: scenarioInput.Test, GetBlockResponse: res}, nil
	}

	return nil, errors.New("scenario type unrecognized")
}
