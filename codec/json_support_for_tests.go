// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strconv"
)

// this file is only required for the json-based contract test in /test/codec

const ISO_DATE_FORMAT = "2006-01-02T15:04:05.000Z07:00"

func (r *SendTransactionRequest) UnmarshalJSON(data []byte) error {
	type OtherFields SendTransactionRequest
	aux := &struct {
		ProtocolVersion     string
		VirtualChainId      string
		InputArguments      []string
		InputArgumentsTypes []string
		*OtherFields
	}{
		OtherFields: (*OtherFields)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	protocolVersion, err := strconv.ParseUint(aux.ProtocolVersion, 10, 32)
	if err != nil {
		panic(err)
	}
	r.ProtocolVersion = uint32(protocolVersion)
	virtualChainId, err := strconv.ParseUint(aux.VirtualChainId, 10, 32)
	if err != nil {
		panic(err)
	}
	r.VirtualChainId = uint32(virtualChainId)
	r.InputArguments = jsonUnmarshalArguments(aux.InputArguments, aux.InputArgumentsTypes)
	return nil
}

func (r *SendTransactionResponse) MarshalJSON() ([]byte, error) {
	type OtherFields SendTransactionResponse
	args, argTypes := jsonMarshalArguments(r.OutputArguments)
	return json.Marshal(&struct {
		BlockHeight          string
		BlockTimestamp       string
		OutputArguments      []string
		OutputArgumentsTypes []string
		OutputEvents         []*jsonEvent
		*OtherFields
	}{
		BlockHeight:          strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:       r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments:      args,
		OutputArgumentsTypes: argTypes,
		OutputEvents:         jsonMarshalEvents(r.OutputEvents),
		OtherFields:          (*OtherFields)(r),
	})
}

func (r *RunQueryRequest) UnmarshalJSON(data []byte) error {
	type OtherFields RunQueryRequest
	aux := &struct {
		ProtocolVersion     string
		VirtualChainId      string
		InputArguments      []string
		InputArgumentsTypes []string
		*OtherFields
	}{
		OtherFields: (*OtherFields)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	protocolVersion, err := strconv.ParseUint(aux.ProtocolVersion, 10, 32)
	if err != nil {
		panic(err)
	}
	r.ProtocolVersion = uint32(protocolVersion)
	virtualChainId, err := strconv.ParseUint(aux.VirtualChainId, 10, 32)
	if err != nil {
		panic(err)
	}
	r.VirtualChainId = uint32(virtualChainId)
	r.InputArguments = jsonUnmarshalArguments(aux.InputArguments, aux.InputArgumentsTypes)
	return nil
}

func (r *RunQueryResponse) MarshalJSON() ([]byte, error) {
	type OtherFields RunQueryResponse
	args, argTypes := jsonMarshalArguments(r.OutputArguments)
	return json.Marshal(&struct {
		BlockHeight          string
		BlockTimestamp       string
		OutputArguments      []string
		OutputArgumentsTypes []string
		OutputEvents         []*jsonEvent
		*OtherFields
	}{
		BlockHeight:          strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:       r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments:      args,
		OutputArgumentsTypes: argTypes,
		OutputEvents:         jsonMarshalEvents(r.OutputEvents),
		OtherFields:          (*OtherFields)(r),
	})
}

func (r *GetTransactionStatusRequest) UnmarshalJSON(data []byte) error {
	type OtherFields GetTransactionStatusRequest
	aux := &struct {
		ProtocolVersion string
		VirtualChainId  string
		*OtherFields
	}{
		OtherFields: (*OtherFields)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	protocolVersion, err := strconv.ParseUint(aux.ProtocolVersion, 10, 32)
	if err != nil {
		panic(err)
	}
	r.ProtocolVersion = uint32(protocolVersion)
	virtualChainId, err := strconv.ParseUint(aux.VirtualChainId, 10, 32)
	if err != nil {
		panic(err)
	}
	r.VirtualChainId = uint32(virtualChainId)
	return nil
}

func (r *GetTransactionStatusResponse) MarshalJSON() ([]byte, error) {
	type OtherFields GetTransactionStatusResponse
	args, argTypes := jsonMarshalArguments(r.OutputArguments)
	return json.Marshal(&struct {
		BlockHeight          string
		BlockTimestamp       string
		OutputArguments      []string
		OutputArgumentsTypes []string
		OutputEvents         []*jsonEvent
		*OtherFields
	}{
		BlockHeight:          strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:       r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments:      args,
		OutputArgumentsTypes: argTypes,
		OutputEvents:         jsonMarshalEvents(r.OutputEvents),
		OtherFields:          (*OtherFields)(r),
	})
}

func (r *GetTransactionReceiptProofRequest) UnmarshalJSON(data []byte) error {
	type OtherFields GetTransactionReceiptProofRequest
	aux := &struct {
		ProtocolVersion string
		VirtualChainId  string
		*OtherFields
	}{
		OtherFields: (*OtherFields)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	protocolVersion, err := strconv.ParseUint(aux.ProtocolVersion, 10, 32)
	if err != nil {
		panic(err)
	}
	r.ProtocolVersion = uint32(protocolVersion)
	virtualChainId, err := strconv.ParseUint(aux.VirtualChainId, 10, 32)
	if err != nil {
		panic(err)
	}
	r.VirtualChainId = uint32(virtualChainId)
	return nil
}

func (r *GetTransactionReceiptProofResponse) MarshalJSON() ([]byte, error) {
	type OtherFields GetTransactionReceiptProofResponse
	args, argTypes := jsonMarshalArguments(r.OutputArguments)
	return json.Marshal(&struct {
		BlockHeight          string
		BlockTimestamp       string
		OutputArguments      []string
		OutputArgumentsTypes []string
		OutputEvents         []*jsonEvent
		*OtherFields
	}{
		BlockHeight:          strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:       r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments:      args,
		OutputArgumentsTypes: argTypes,
		OutputEvents:         jsonMarshalEvents(r.OutputEvents),
		OtherFields:          (*OtherFields)(r),
	})
}

func (r *GetBlockRequest) UnmarshalJSON(data []byte) error {
	type OtherFields GetBlockRequest
	aux := &struct {
		ProtocolVersion string
		VirtualChainId  string
		BlockHeight     string
		*OtherFields
	}{
		OtherFields: (*OtherFields)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	protocolVersion, err := strconv.ParseUint(aux.ProtocolVersion, 10, 32)
	if err != nil {
		panic(err)
	}
	r.ProtocolVersion = uint32(protocolVersion)
	virtualChainId, err := strconv.ParseUint(aux.VirtualChainId, 10, 32)
	if err != nil {
		panic(err)
	}
	r.VirtualChainId = uint32(virtualChainId)
	blockHeight, err := strconv.ParseUint(aux.BlockHeight, 10, 64)
	if err != nil {
		panic(err)
	}
	r.BlockHeight = uint64(blockHeight)
	return nil
}

func (r *GetBlockResponse) MarshalJSON() ([]byte, error) {
	type OtherFields GetBlockResponse
	return json.Marshal(&struct {
		BlockHeight    string
		BlockTimestamp string
		*OtherFields
	}{
		BlockHeight:    strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp: r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OtherFields:    (*OtherFields)(r),
	})
}

func (r *TransactionsBlockHeader) MarshalJSON() ([]byte, error) {
	type OtherFields TransactionsBlockHeader
	return json.Marshal(&struct {
		ProtocolVersion string
		VirtualChainId  string
		BlockHeight     string
		Timestamp       string
		NumTransactions string
		*OtherFields
	}{
		ProtocolVersion: strconv.FormatUint(uint64(r.ProtocolVersion), 10),
		VirtualChainId:  strconv.FormatUint(uint64(r.VirtualChainId), 10),
		BlockHeight:     strconv.FormatUint(r.BlockHeight, 10),
		Timestamp:       r.Timestamp.UTC().Format(ISO_DATE_FORMAT),
		NumTransactions: strconv.FormatUint(uint64(r.NumTransactions), 10),
		OtherFields:     (*OtherFields)(r),
	})
}

func (r *ResultsBlockHeader) MarshalJSON() ([]byte, error) {
	type OtherFields ResultsBlockHeader
	return json.Marshal(&struct {
		ProtocolVersion        string
		VirtualChainId         string
		BlockHeight            string
		Timestamp              string
		NumTransactionReceipts string
		*OtherFields
	}{
		ProtocolVersion:        strconv.FormatUint(uint64(r.ProtocolVersion), 10),
		VirtualChainId:         strconv.FormatUint(uint64(r.VirtualChainId), 10),
		BlockHeight:            strconv.FormatUint(r.BlockHeight, 10),
		Timestamp:              r.Timestamp.UTC().Format(ISO_DATE_FORMAT),
		NumTransactionReceipts: strconv.FormatUint(uint64(r.NumTransactionReceipts), 10),
		OtherFields:            (*OtherFields)(r),
	})
}

func (r *BlockTransaction) MarshalJSON() ([]byte, error) {
	type OtherFields BlockTransaction
	inArgs, inArgTypes := jsonMarshalArguments(r.InputArguments)
	outArgs, outArgTypes := jsonMarshalArguments(r.OutputArguments)
	return json.Marshal(&struct {
		ProtocolVersion      string
		VirtualChainId       string
		Timestamp            string
		InputArguments       []string
		InputArgumentsTypes  []string
		OutputArguments      []string
		OutputArgumentsTypes []string
		OutputEvents         []*jsonEvent
		*OtherFields
	}{
		ProtocolVersion:      strconv.FormatUint(uint64(r.ProtocolVersion), 10),
		VirtualChainId:       strconv.FormatUint(uint64(r.VirtualChainId), 10),
		Timestamp:            r.Timestamp.UTC().Format(ISO_DATE_FORMAT),
		InputArguments:       inArgs,
		InputArgumentsTypes:  inArgTypes,
		OutputArguments:      outArgs,
		OutputArgumentsTypes: outArgTypes,
		OutputEvents:         jsonMarshalEvents(r.OutputEvents),
		OtherFields:          (*OtherFields)(r),
	})
}

func jsonUnmarshalArguments(arguments []string, argumentsTypes []string) []interface{} {
	res := []interface{}{}
	for index, arg := range arguments {
		if len(argumentsTypes) > index {
			switch argumentsTypes[index] {
			case "uint32":
				num, err := strconv.ParseInt(arg, 10, 64)
				if err != nil {
					panic(err)
				}
				res = append(res, uint32(num))
			case "uint64":
				num, err := strconv.ParseInt(arg, 10, 64)
				if err != nil {
					panic(err)
				}
				res = append(res, uint64(num))
			case "string":
				res = append(res, arg)
			case "bytes":
				bytes, err := hex.DecodeString(arg)
				if err != nil {
					panic(err)
				}
				res = append(res, bytes)
			case "bool":
				res = append(res, arg == "1")
			case "uint256":
				bytes, err := hex.DecodeString(arg)
				if err != nil {
					panic(err)
				}
				obj := big.NewInt(0)
				obj.SetBytes(bytes)
				res = append(res, obj)
			case "bytes20":
				bytes, err := hex.DecodeString(arg)
				if err != nil {
					panic(err)
				}
				bytes20 := [20]byte{}
				copy(bytes20[:], bytes)
				res = append(res, bytes20)
			case "bytes32":
				bytes, err := hex.DecodeString(arg)
				if err != nil {
					panic(err)
				}
				bytes32 := [32]byte{}
				copy(bytes32[:], bytes)
				res = append(res, bytes32)
			default:
				res = append(res, arg)
			}
		} else {
			res = append(res, arg)
		}
	}
	return res
}

func jsonMarshalArguments(arguments []interface{}) ([]string, []string) {
	res := []string{}
	resTypes := []string{}
	for _, arg := range arguments {
		switch arg.(type) {
		case uint32:
			res = append(res, strconv.FormatUint(uint64(arg.(uint32)), 10))
			resTypes = append(resTypes, "uint32")
		case uint64:
			res = append(res, strconv.FormatUint(uint64(arg.(uint64)), 10))
			resTypes = append(resTypes, "uint64")
		case string:
			res = append(res, arg.(string))
			resTypes = append(resTypes, "string")
		case []byte:
			res = append(res, hex.EncodeToString(arg.([]byte)))
			resTypes = append(resTypes, "bytes")
		case bool:
			if arg.(bool) {
				res = append(res, "1")
			} else {
				res = append(res, "0")
			}
			resTypes = append(resTypes, "bool")
		case *big.Int:
			actual := [32]byte{}
			b := arg.(*big.Int).Bytes()
			copy(actual[32-len(b):], b)
			res = append(res, hex.EncodeToString(actual[:]))
			resTypes = append(resTypes, "uint256")
		case [20]byte:
			obj := arg.([20]byte)
			res = append(res, hex.EncodeToString(obj[:]))
			resTypes = append(resTypes, "bytes20")
		case [32]byte:
			obj := arg.([32]byte)
			res = append(res, hex.EncodeToString(obj[:]))
			resTypes = append(resTypes, "bytes32")
		default:
			panic("unsupported type in json marshal of method arguments")
		}
	}
	return res, resTypes
}

type jsonEvent struct {
	ContractName   string
	EventName      string
	Arguments      []string
	ArgumentsTypes []string
}

func jsonMarshalEvents(events []*Event) []*jsonEvent {
	res := []*jsonEvent{}
	for _, event := range events {
		args, argTypes := jsonMarshalArguments(event.Arguments)
		res = append(res, &jsonEvent{
			ContractName:   event.ContractName,
			EventName:      event.EventName,
			Arguments:      args,
			ArgumentsTypes: argTypes,
		})
	}
	return res
}

func jsonUnmarshalEvents(events []*jsonEvent) []*Event {
	res := []*Event{}
	for _, event := range events {
		res = append(res, &Event{
			ContractName: event.ContractName,
			EventName:    event.EventName,
			Arguments:    jsonUnmarshalArguments(event.Arguments, event.ArgumentsTypes),
		})
	}
	return res
}
