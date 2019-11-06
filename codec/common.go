package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"time"
)

type Response struct {
	RequestStatus   RequestStatus
	BlockHeight     uint64
	BlockTimestamp  time.Time
}

type ReadResponse struct {
	Response
	ExecutionResult ExecutionResult
	OutputArguments []interface{}
	OutputEvents    []*Event
}

type TransactionResponse struct {
	ReadResponse
	TransactionStatus TransactionStatus
	TxHash            []byte
}

type requestResponser interface {
	RequestResult() *client.RequestResult

}

type txResponse interface {
	requestResponser
	TransactionReceipt() *protocol.TransactionReceipt
}

func NewTransactionResponse(res txResponse, outputArgumentArray []interface{}, outputEventArray []*Event, executionResult ExecutionResult, requestStatus RequestStatus, transactionStatus TransactionStatus) TransactionResponse {
	return TransactionResponse{
		TxHash:            res.TransactionReceipt().Txhash(),
		TransactionStatus: transactionStatus,
		ReadResponse: NewReadResponse(res, outputArgumentArray, outputEventArray, executionResult, requestStatus),
	}
}

func NewReadResponse(res requestResponser, outputArgumentArray []interface{}, outputEventArray []*Event, executionResult ExecutionResult, requestStatus RequestStatus) ReadResponse {
	return ReadResponse{
		OutputArguments: outputArgumentArray,
		OutputEvents:    outputEventArray,
		ExecutionResult: executionResult,
		Response:        NewResponse(res, requestStatus),
	}
}

func NewResponse(res requestResponser, requestStatus RequestStatus) Response {
	return Response{
		BlockHeight:    uint64(res.RequestResult().BlockHeight()),
		BlockTimestamp: time.Unix(0, int64(res.RequestResult().BlockTimestamp())),
		RequestStatus:  requestStatus,
	}
}
