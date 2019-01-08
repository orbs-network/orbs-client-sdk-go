package jsoncodec

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"strconv"
)

type Arg struct {
	Type  string
	Value string
}

func UnmarshalArgs(args []*Arg, getTestKeyFromFile func(string) *RawKey) ([]interface{}, error) {
	res := []interface{}{}
	for i, arg := range args {
		switch arg.Type {
		case "uint32":
			val, err := strconv.ParseUint(arg.Value, 10, 32)
			if err != nil {
				return nil, errors.Errorf("Value of argument %d should be a string containing the numeric value\n\nCurrent value: '%s'", i, arg.Value)
			}
			res = append(res, uint32(val))
		case "uint64":
			val, err := strconv.ParseUint(arg.Value, 10, 64)
			if err != nil {
				return nil, errors.Errorf("Value of argument %d should be a string containing the numeric value\n\nCurrent value: '%s'", i, arg.Value)
			}
			res = append(res, uint64(val))
		case "string":
			res = append(res, string(arg.Value))
		case "bytes":
			val, err := hex.DecodeString(arg.Value)
			if err != nil {
				return nil, errors.Errorf("Value of argument %d should be a string containing the bytes in hex\n\nCurrent value: '%s'", i, arg.Value)
			}
			res = append(res, []byte(val))
		case "gamma:keys-file-address":
			key := getTestKeyFromFile(arg.Value)
			res = append(res, []byte(key.Address))
		default:
			supported := "Supported types are: uint32 uint64 string bytes gamma:keys-file-address"
			return nil, errors.Errorf("Type of argument %d '%s' is unsupported\n\n%s", i, arg.Type, supported)
		}
	}
	return res, nil
}

func MarshalArgs(arguments []interface{}) []*Arg {
	res := []*Arg{}
	for _, arg := range arguments {
		switch arg.(type) {
		case uint32:
			res = append(res, &Arg{"uint32", strconv.FormatUint(uint64(arg.(uint32)), 10)})
		case uint64:
			res = append(res, &Arg{"uint64", strconv.FormatUint(uint64(arg.(uint64)), 10)})
		case string:
			res = append(res, &Arg{"string", arg.(string)})
		case []byte:
			res = append(res, &Arg{"bytes", hex.EncodeToString(arg.([]byte))})
		default:
			panic("unsupported type in json marshal of method arguments")
		}
	}
	return res
}
