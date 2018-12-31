package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
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
		default:
			err = errors.Errorf("given method argument %d has unsupported type (%T), supported: (uint32) (uint64) (string) ([]byte)", index, arg)
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
		default:
			err = errors.Errorf("received method argument %d has unknown type: %s", index, argument.StringType())
			return
		}
		index++
	}
	return
}
