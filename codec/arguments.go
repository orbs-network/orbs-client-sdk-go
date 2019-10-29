// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
	"math/big"
)

func argumentsBuilders(args []interface{}) (res []*protocol.ArgumentBuilder, err error) {
	res = []*protocol.ArgumentBuilder{}
	for index, arg := range args {
		switch arg.(type) {
		case uint32:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_UINT_32_VALUE, Uint32Value: arg.(uint32)})
		case uint64:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_UINT_64_VALUE, Uint64Value: arg.(uint64)})
		case string:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_STRING_VALUE, StringValue: arg.(string)})
		case []byte:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_BYTES_VALUE, BytesValue: arg.([]byte)})
		case bool:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_BOOL_VALUE, BoolValue: arg.(bool)})
		case *big.Int:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_UINT_256_VALUE, Uint256Value: arg.(*big.Int)})
		case [20]byte:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_BYTES_20_VALUE, Bytes20Value: arg.([20]byte)})
		case [32]byte:
			res = append(res, &protocol.ArgumentBuilder{Type: protocol.ARGUMENT_TYPE_BYTES_32_VALUE, Bytes32Value: arg.([32]byte)})
		default:
			err = errors.Errorf("given method argument %d has unsupported type (%T), supported: (uint32) (uint64) (string) ([]byte) (bool) (uint256) ([20]byte) ([32]byte)", index, arg)
			return
		}
	}
	return
}

func argumentsArray(args []interface{}) (*protocol.ArgumentArray, error) {
	builders, err := argumentsBuilders(args)
	if err != nil {
		return nil, err
	}
	return (&protocol.ArgumentArrayBuilder{Arguments: builders}).Build(), nil
}

func PackedArgumentsEncode(args []interface{}) ([]byte, error) {
	argArray, err := argumentsArray(args)
	if err != nil {
		return nil, err
	}
	return argArray.RawArgumentsArray(), nil
}

func PackedArgumentsDecode(buf []byte) (res []interface{}, err error) {
	res = []interface{}{}
	argsArray := protocol.ArgumentArrayReader(buf)
	index := 0
	for i := argsArray.ArgumentsIterator(); i.HasNext(); {
		argument := i.NextArguments()
		switch argument.Type() {
		case protocol.ARGUMENT_TYPE_UINT_32_VALUE:
			res = append(res, argument.Uint32Value())
		case protocol.ARGUMENT_TYPE_UINT_64_VALUE:
			res = append(res, argument.Uint64Value())
		case protocol.ARGUMENT_TYPE_STRING_VALUE:
			res = append(res, argument.StringValue())
		case protocol.ARGUMENT_TYPE_BYTES_VALUE:
			res = append(res, argument.BytesValue())
		case protocol.ARGUMENT_TYPE_BOOL_VALUE:
			res = append(res, argument.BoolValue())
		case protocol.ARGUMENT_TYPE_UINT_256_VALUE:
			res = append(res, argument.Uint256Value())
		case protocol.ARGUMENT_TYPE_BYTES_20_VALUE:
			res = append(res, argument.Bytes20Value())
		case protocol.ARGUMENT_TYPE_BYTES_32_VALUE:
			res = append(res, argument.Bytes32Value())
		default:
			err = errors.Errorf("received method argument %d has unknown type: %s", index, argument.StringType())
			return
		}
		index++
	}
	return
}
