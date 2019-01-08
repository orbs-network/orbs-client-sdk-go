package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

type RequestStatus string

const (
	REQUEST_STATUS_PARSE_ERROR  RequestStatus = "<PARSE_ERROR>"
	REQUEST_STATUS_COMPLETED    RequestStatus = "COMPLETED"
	REQUEST_STATUS_IN_PROCESS   RequestStatus = "IN_PROCESS"
	REQUEST_STATUS_BAD_REQUEST  RequestStatus = "BAD_REQUEST"
	REQUEST_STATUS_CONGESTION   RequestStatus = "CONGESTION"
	REQUEST_STATUS_SYSTEM_ERROR RequestStatus = "SYSTEM_ERROR"
)

func (x RequestStatus) String() string {
	return string(x)
}

func requestStatusDecode(requestStatus protocol.RequestStatus) (RequestStatus, error) {
	switch requestStatus {
	case protocol.REQUEST_STATUS_RESERVED:
		return REQUEST_STATUS_PARSE_ERROR, errors.Errorf("reserved RequestStatus received")
	case protocol.REQUEST_STATUS_COMPLETED:
		return REQUEST_STATUS_COMPLETED, nil
	case protocol.REQUEST_STATUS_IN_PROCESS:
		return REQUEST_STATUS_IN_PROCESS, nil
	case protocol.REQUEST_STATUS_BAD_REQUEST:
		return REQUEST_STATUS_BAD_REQUEST, nil
	case protocol.REQUEST_STATUS_CONGESTION:
		return REQUEST_STATUS_CONGESTION, nil
	case protocol.REQUEST_STATUS_SYSTEM_ERROR:
		return REQUEST_STATUS_SYSTEM_ERROR, nil
	default:
		return REQUEST_STATUS_PARSE_ERROR, errors.Errorf("unsupported RequestStatus received: %d", requestStatus)
	}
}
