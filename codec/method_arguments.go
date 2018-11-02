package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

func methodArgumentsBuilders(args []interface{}) (res []*protocol.MethodArgumentBuilder, err error) {
	res = []*protocol.MethodArgumentBuilder{}
	for index, arg := range args {
		switch arg.(type) {
		case uint32:
			res = append(res, &protocol.MethodArgumentBuilder{Name: "uint32", Type: protocol.METHOD_ARGUMENT_TYPE_UINT_32_VALUE, Uint32Value: arg.(uint32)})
		case uint64:
			res = append(res, &protocol.MethodArgumentBuilder{Name: "uint64", Type: protocol.METHOD_ARGUMENT_TYPE_UINT_64_VALUE, Uint64Value: arg.(uint64)})
		case string:
			res = append(res, &protocol.MethodArgumentBuilder{Name: "string", Type: protocol.METHOD_ARGUMENT_TYPE_STRING_VALUE, StringValue: arg.(string)})
		case []byte:
			res = append(res, &protocol.MethodArgumentBuilder{Name: "bytes", Type: protocol.METHOD_ARGUMENT_TYPE_BYTES_VALUE, BytesValue: arg.([]byte)})
		default:
			err = errors.Errorf("given method argument %d has unsupported type (%T), supported: (uint32) (uint64) (string) ([]byte)", index, arg)
			return
		}
	}
	return
}

func methodArgumentsArray(args []interface{}) (*protocol.MethodArgumentArray, error) {
	builders, err := methodArgumentsBuilders(args)
	if err != nil {
		return nil, err
	}
	return (&protocol.MethodArgumentArrayBuilder{Arguments: builders}).Build(), nil
}

func MethodArgumentsOpaqueEncode(args []interface{}) ([]byte, error) {
	argArray, err := methodArgumentsArray(args)
	if err != nil {
		return nil, err
	}
	return argArray.RawArgumentsArray(), nil
}

func MethodArgumentsOpaqueDecode(buf []byte) (res []interface{}, err error) {
	res = []interface{}{}
	argsArray := protocol.MethodArgumentArrayReader(buf)
	index := 0
	for i := argsArray.ArgumentsIterator(); i.HasNext(); {
		methodArgument := i.NextArguments()
		switch methodArgument.Type() {
		case protocol.METHOD_ARGUMENT_TYPE_UINT_32_VALUE:
			res = append(res, methodArgument.Uint32Value())
		case protocol.METHOD_ARGUMENT_TYPE_UINT_64_VALUE:
			res = append(res, methodArgument.Uint64Value())
		case protocol.METHOD_ARGUMENT_TYPE_STRING_VALUE:
			res = append(res, methodArgument.StringValue())
		case protocol.METHOD_ARGUMENT_TYPE_BYTES_VALUE:
			res = append(res, methodArgument.BytesValue())
		default:
			err = errors.Errorf("received method argument %d has unknown type: %s", index, methodArgument.StringType())
			return
		}
		index++
	}
	return
}
