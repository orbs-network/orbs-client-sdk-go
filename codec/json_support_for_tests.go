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
				bytes, err := base64.StdEncoding.DecodeString(arg)
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
			res = append(res, base64.StdEncoding.EncodeToString(arg.([]byte)))
			resTypes = append(resTypes, "bytes")
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
