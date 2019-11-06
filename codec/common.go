package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"time"
)

type Response struct {
	RequestStatus  RequestStatus
	BlockHeight    uint64
	BlockTimestamp time.Time
}

type ReadResponse struct {
	*Response
	ExecutionResult ExecutionResult
	OutputArguments []interface{}
	OutputEvents    []*Event
}

type TransactionResponse struct {
	*ReadResponse
	TransactionStatus TransactionStatus
	TxHash            []byte
}

type ReceiptLike interface {
	ExecutionResult() protocol.ExecutionResult
	RawOutputArgumentArrayWithHeader() []byte
	RawOutputEventsArrayWithHeader() []byte
}

type requestResponser interface {
	RequestResult() *client.RequestResult
}

type txResponse interface {
	requestResponser
	TransactionStatus() protocol.TransactionStatus
	TransactionReceipt() *protocol.TransactionReceipt
	RawTransactionReceipt() []byte
}

func NewTransactionResponse(res txResponse) (*TransactionResponse, error) {
	readResponse, err := NewReadResponse(res, res.RawTransactionReceipt(), res.TransactionReceipt())
	if err != nil {
		return nil, err
	}

	transactionStatus, err := transactionStatusDecode(res.TransactionStatus())
	if err != nil {
		return nil, err
	}

	return &TransactionResponse{
		TxHash:            res.TransactionReceipt().Txhash(),
		TransactionStatus: transactionStatus,
		ReadResponse:      readResponse,
	}, nil
}

func NewReadResponse(res requestResponser, rawReceipt []byte, receipt ReceiptLike) (*ReadResponse, error) {
	response, err := NewResponse(res)
	if err != nil {
		return nil, err
	}

	executionResult := EXECUTION_RESULT_NOT_EXECUTED
	if len(rawReceipt) > 0 {
		executionResult, err = executionResultDecode(receipt.ExecutionResult())
		if err != nil {
			return nil, err
		}
	}

	outputArgumentArray, err := PackedArgumentsDecode(receipt.RawOutputArgumentArrayWithHeader())
	if err != nil {
		return nil, err
	}

	outputEventArray, err := PackedEventsDecode(receipt.RawOutputEventsArrayWithHeader())
	if err != nil {
		return nil, err
	}

	return &ReadResponse{
		OutputArguments: outputArgumentArray,
		OutputEvents:    outputEventArray,
		ExecutionResult: executionResult,
		Response:        response,
	}, nil
}

func NewResponse(res requestResponser) (*Response, error) {
	requestStatus, err := requestStatusDecode(res.RequestResult().RequestStatus())
	if err != nil {
		return nil, err
	}

	return &Response{
		BlockHeight:    uint64(res.RequestResult().BlockHeight()),
		BlockTimestamp: time.Unix(0, int64(res.RequestResult().BlockTimestamp())),
		RequestStatus:  requestStatus,
	}, nil
}
