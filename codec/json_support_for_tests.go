package codec

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

// this file is only required for the json-based contract test in /test/codec

const ISO_DATE_FORMAT = "2006-01-02T15:04:05.000Z07:00"

func (r *SendTransactionRequest) UnmarshalJSON(data []byte) error {
	type OtherFields SendTransactionRequest
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

func (r *SendTransactionRequest) FixInputArgumentTypes(argumentsTypes []string) {
	r.InputArguments = jsonUnmarshalMethodArguments(r.InputArguments, argumentsTypes)
}

func (r *SendTransactionResponse) MarshalJSON() ([]byte, error) {
	type OtherFields SendTransactionResponse
	return json.Marshal(&struct {
		BlockHeight     string
		BlockTimestamp  string
		OutputArguments []string
		*OtherFields
	}{
		BlockHeight:     strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:  r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments: jsonMarshalMethodArguments(r.OutputArguments),
		OtherFields:     (*OtherFields)(r),
	})
}

func (r *CallMethodRequest) UnmarshalJSON(data []byte) error {
	type OtherFields CallMethodRequest
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

func (r *CallMethodRequest) FixInputArgumentTypes(argumentsTypes []string) {
	r.InputArguments = jsonUnmarshalMethodArguments(r.InputArguments, argumentsTypes)
}

func (r *CallMethodResponse) MarshalJSON() ([]byte, error) {
	type OtherFields CallMethodResponse
	return json.Marshal(&struct {
		BlockHeight     string
		BlockTimestamp  string
		OutputArguments []string
		*OtherFields
	}{
		BlockHeight:     strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:  r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments: jsonMarshalMethodArguments(r.OutputArguments),
		OtherFields:     (*OtherFields)(r),
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
	return json.Marshal(&struct {
		BlockHeight     string
		BlockTimestamp  string
		OutputArguments []string
		*OtherFields
	}{
		BlockHeight:     strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:  r.BlockTimestamp.UTC().Format(ISO_DATE_FORMAT),
		OutputArguments: jsonMarshalMethodArguments(r.OutputArguments),
		OtherFields:     (*OtherFields)(r),
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

func jsonUnmarshalMethodArguments(arguments []interface{}, argumentsTypes []string) []interface{} {
	res := []interface{}{}
	for index, arg := range arguments {
		if len(argumentsTypes) > index {
			switch argumentsTypes[index] {
			case "uint32":
				num, err := strconv.ParseInt(arg.(string), 10, 64)
				if err != nil {
					panic(err)
				}
				res = append(res, uint32(num))
			case "uint64":
				num, err := strconv.ParseInt(arg.(string), 10, 64)
				if err != nil {
					panic(err)
				}
				res = append(res, uint64(num))
			case "string":
				res = append(res, arg.(string))
			case "bytes":
				bytes, err := base64.StdEncoding.DecodeString(arg.(string))
				if err != nil {
					panic(err)
				}
				res = append(res, bytes)
			default:
				res = append(res, arg)
			}
		} else {
			res = append(res, arg)
		}
	}
	return res
}

func jsonMarshalMethodArguments(arguments []interface{}) []string {
	res := []string{}
	for _, arg := range arguments {
		switch arg.(type) {
		case uint32:
			res = append(res, strconv.FormatUint(uint64(arg.(uint32)), 10))
		case uint64:
			res = append(res, strconv.FormatUint(uint64(arg.(uint64)), 10))
		case string:
			res = append(res, arg.(string))
		case []byte:
			res = append(res, base64.StdEncoding.EncodeToString(arg.([]byte)))
		default:
			panic("unsupported type in json marshal of method arguments")
		}
	}
	return res
}
