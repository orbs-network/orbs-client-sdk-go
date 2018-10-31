package codec

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/digest"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/pkg/errors"
	"time"
)

type GetTransactionStatusRequest struct {
	TxId []byte
}

type GetTransactionStatusResponse struct {
	RequestStatus     RequestStatus
	TxHash            []byte
	ExecutionResult   ExecutionResult
	OutputArguments   []interface{}
	TransactionStatus TransactionStatus
	BlockHeight       uint64
	BlockTimestamp    time.Time
}

func EncodeGetTransactionStatusRequest(req *GetTransactionStatusRequest) ([]byte, error) {
	// validate
	if len(req.TxId) != digest.TX_ID_SIZE_BYTES {
		return nil, errors.Errorf("expected TxId length %d, %d given", digest.TX_ID_SIZE_BYTES, len(req.TxId))
	}

	// extract txid
	txHash, txTimestamp, err := digest.ExtractTxId(req.TxId)
	if err != nil {
		return nil, err
	}

	// encode request
	res := (&client.GetTransactionStatusRequestBuilder{
		TransactionTimestamp: txTimestamp,
		Txhash:               primitives.Sha256(txHash),
	}).Build()

	// return
	return res.Raw(), nil
}

func DecodeGetTransactionStatusResponse(buf []byte) (*GetTransactionStatusResponse, error) {
	// decode response
	res := client.GetTransactionStatusResponseReader(buf)
	if !res.IsValid() {
		return nil, errors.New("response is corrupt and cannot be decoded")
	}

	// decode request status
	requestStatus, err := requestStatusDecode(res.RequestStatus())
	if err != nil {
		return nil, err
	}

	// decode execution result
	executionResult, err := executionResultDecode(res.TransactionReceipt().ExecutionResult())
	if err != nil {
		return nil, err
	}

	// decode method arguments
	outputArgumentArray, err := MethodArgumentsOpaqueDecode(res.TransactionReceipt().RawOutputArgumentArrayWithHeader())
	if err != nil {
		return nil, err
	}

	// decode transaction status
	transactionStatus, err := transactionStatusDecode(res.TransactionStatus())
	if err != nil {
		return nil, err
	}

	// return
	return &GetTransactionStatusResponse{
		RequestStatus:     requestStatus,
		TxHash:            res.TransactionReceipt().Txhash(),
		ExecutionResult:   executionResult,
		OutputArguments:   outputArgumentArray,
		TransactionStatus: transactionStatus,
		BlockHeight:       uint64(res.BlockHeight()),
		BlockTimestamp:    time.Unix(0, int64(res.BlockTimestamp())),
	}, nil
}
