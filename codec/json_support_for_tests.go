// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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

func jsonPreUnmarshalForArray(arg string) []string {
	var res []string
	if err := json.Unmarshal([]byte(arg), &res); err != nil {
		panic(fmt.Sprintf("parse of string contain array of string %s\n", err.Error()))
	}
	return res
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
			case "uint32Array":
				argArr := jsonPreUnmarshalForArray(arg)
				var arr []uint32
				for _, internalArg := range argArr {
					num, err := strconv.ParseInt(internalArg, 10, 64)
					if err != nil {
						panic(err)
					}
					arr = append(arr, uint32(num))
				}
				res = append(res, arr)
			case "uint64Array":
				argArr := jsonPreUnmarshalForArray(arg)
				//				fmt.Printf("uints str %v %T\n", argArr, argArr)
				var arr []uint64
				for _, internalArg := range argArr {
					//					fmt.Printf("uint str %v %T\n", internalArg, internalArg)
					num, err := strconv.ParseInt(internalArg, 10, 64)
					if err != nil {
						panic(err)
					}
					//					fmt.Printf("uint num %v %T\n", num, num)
					arr = append(arr, uint64(num))
				}
				res = append(res, arr)
			case "stringArray":
				argArr := jsonPreUnmarshalForArray(arg)
				res = append(res, argArr)
			case "bytesArray":
				argArr := jsonPreUnmarshalForArray(arg)
				var arr [][]byte
				for _, internalArg := range argArr {
					bytes, err := hex.DecodeString(internalArg)
					if err != nil {
						panic(err)
					}
					arr = append(arr, bytes)
				}
				res = append(res, arr)
			case "boolArray":
				argArr := jsonPreUnmarshalForArray(arg)
				//				fmt.Printf("bools str %v %T\n", argArr, argArr)
				var arr []bool
				for _, internalArg := range argArr {
					//					fmt.Printf("bools str %v %T\n", internalArg, internalArg)
					arr = append(arr, internalArg == "1")
				}
				res = append(res, arr)
			case "uint256Array":
				argArr := jsonPreUnmarshalForArray(arg)
				var arr []*big.Int
				for _, internalArg := range argArr {
					bytes, err := hex.DecodeString(internalArg)
					if err != nil {
						panic(err)
					}
					obj := big.NewInt(0)
					obj.SetBytes(bytes)
					arr = append(arr, obj)
				}
				res = append(res, arr)
			case "bytes20Array":
				argArr := jsonPreUnmarshalForArray(arg)
				var arr [][20]byte
				for _, internalArg := range argArr {
					bytes, err := hex.DecodeString(internalArg)
					if err != nil {
						panic(err)
					}
					bytes20 := [20]byte{}
					copy(bytes20[:], bytes)
					arr = append(arr, bytes20)
				}
				res = append(res, arr)
			case "bytes32Array":
				argArr := jsonPreUnmarshalForArray(arg)
				var arr [][32]byte
				for _, internalArg := range argArr {
					bytes, err := hex.DecodeString(internalArg)
					if err != nil {
						panic(err)
					}
					bytes32 := [32]byte{}
					copy(bytes32[:], bytes)
					arr = append(arr, bytes32)
				}
				res = append(res, arr)
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
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "uint32")
		case uint64:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "uint64")
		case string:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "string")
		case []byte:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "bytes")
		case bool:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "bool")
		case *big.Int:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "uint256")
		case [20]byte:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "bytes20")
		case [32]byte:
			res = append(res, jsonMarshalPrimitiveScalar(arg))
			resTypes = append(resTypes, "bytes32")
		case []uint32:
			internalArgs := arg.([]uint32)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "uint32Array")
		case []uint64:
			internalArgs := arg.([]uint64)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "uint64Array")
		case []string:
			internalArgs := arg.([]string)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "stringArray")
		case [][]byte:
			internalArgs := arg.([][]byte)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "bytesArray")
		case []bool:
			internalArgs := arg.([]bool)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "boolArray")
		case []*big.Int:
			internalArgs := arg.([]*big.Int)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "uint256Array")
		case [][20]byte:
			internalArgs := arg.([][20]byte)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "bytes20Array")
		case [][32]byte:
			internalArgs := arg.([][32]byte)
			var resArgs []string
			for _, v := range internalArgs {
				resArgs = append(resArgs, jsonMarshalPrimitiveScalar(v))
			}
			res = append(res, jsonMarshalArray(resArgs))
			resTypes = append(resTypes, "bytes32Array")
		default:
			panic("unsupported type in json marshal of method arguments")
		}
	}
	return res, resTypes
}

// can only run inside the bigger function - doesn't check errors
func jsonMarshalPrimitiveScalar(arg interface{}) string {
	switch arg := arg.(type) {
	case uint32:
		return strconv.FormatUint(uint64(arg), 10)
	case uint64:
		return strconv.FormatUint(arg, 10)
	case string:
		return arg
	case []byte:
		return hex.EncodeToString(arg)
	case bool:
		if arg {
			return "1"
		} else {
			return "0"
		}
	case *big.Int:
		actual := [32]byte{}
		b := arg.Bytes()
		copy(actual[32-len(b):], b)
		return hex.EncodeToString(actual[:])
	case [20]byte:
		obj := arg
		return hex.EncodeToString(obj[:])
	case [32]byte:
		obj := arg
		return hex.EncodeToString(obj[:])
	default:
		panic("")
	}
	return ""
}

func jsonMarshalArray(a []string) string {
	bytes, _ := json.Marshal(a)
	return string(bytes)
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
