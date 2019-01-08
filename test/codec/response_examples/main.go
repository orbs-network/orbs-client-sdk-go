package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"math"
	"time"
)

func main() {
	t1, _ := time.Parse(time.RFC3339, "2010-11-21T10:00:00.000Z")
	t2, _ := time.Parse(time.RFC3339, "2020-11-21T20:00:00.000Z")

	h1, _ := hex.DecodeString("cf80cd8aed482d5d1527d7dc72fceff84e6326592848447d2dc0b0e87dfc9a90")

	a1, _ := codec.PackedArgumentsEncode([]interface{}{uint32(1), uint64(2), "hello", []byte{0x01, 0x02, 0x03}})
	a2, _ := codec.PackedArgumentsEncode([]interface{}{})
	a3, _ := codec.PackedArgumentsEncode([]interface{}{uint64(math.MaxUint64 - 1000)})

	e1 := codec.PackedEventsEncode([]*protocol.EventBuilder{
		{ContractName: "Contract1", EventName: "Event1", OutputArgumentArray: a1},
		{ContractName: "Contract2", EventName: "Event2", OutputArgumentArray: a2},
	})
	e2 := codec.PackedEventsEncode([]*protocol.EventBuilder{
		{ContractName: "Contract3", EventName: "Event3", OutputArgumentArray: a3},
	})

	r1 := (&client.SendTransactionResponseBuilder{
		RequestResult: &client.RequestResultBuilder{
			RequestStatus:  protocol.REQUEST_STATUS_COMPLETED,
			BlockHeight:    2135,
			BlockTimestamp: primitives.TimestampNano(t1.UnixNano()),
		},
		TransactionReceipt: &protocol.TransactionReceiptBuilder{
			Txhash:              h1,
			ExecutionResult:     protocol.EXECUTION_RESULT_ERROR_SMART_CONTRACT,
			OutputArgumentArray: a1,
			OutputEventsArray:   e2,
		},
		TransactionStatus: protocol.TRANSACTION_STATUS_COMMITTED,
	}).Build()
	fmt.Printf(`"SendTransactionResponse": "%s"`+"\n\n", base64.StdEncoding.EncodeToString(r1.Raw()))

	r2 := (&client.RunQueryResponseBuilder{
		RequestResult: &client.RequestResultBuilder{
			RequestStatus:  protocol.REQUEST_STATUS_BAD_REQUEST,
			BlockHeight:    87438,
			BlockTimestamp: primitives.TimestampNano(t1.UnixNano()),
		},
		QueryResult: &protocol.QueryResultBuilder{
			ExecutionResult:     protocol.EXECUTION_RESULT_SUCCESS,
			OutputArgumentArray: a2,
			OutputEventsArray:   e1,
		},
	}).Build()
	fmt.Printf(`"RunQueryResponse": "%s"`+"\n\n", base64.StdEncoding.EncodeToString(r2.Raw()))

	r3 := (&client.GetTransactionStatusResponseBuilder{
		RequestResult: &client.RequestResultBuilder{
			RequestStatus:  protocol.REQUEST_STATUS_BAD_REQUEST,
			BlockHeight:    math.MaxUint64 - 5000,
			BlockTimestamp: primitives.TimestampNano(t2.UnixNano()),
		},
		TransactionReceipt: &protocol.TransactionReceiptBuilder{
			Txhash:              h1,
			ExecutionResult:     protocol.EXECUTION_RESULT_ERROR_INPUT,
			OutputArgumentArray: a3,
		},
		TransactionStatus: protocol.TRANSACTION_STATUS_REJECTED_UNKNOWN_SIGNER_SCHEME,
	}).Build()
	fmt.Printf(`"GetTransactionStatusResponse": "%s"`+"\n\n", base64.StdEncoding.EncodeToString(r3.Raw()))

	r4 := (&client.GetTransactionReceiptProofResponseBuilder{
		RequestResult: &client.RequestResultBuilder{
			RequestStatus:  protocol.REQUEST_STATUS_IN_PROCESS,
			BlockHeight:    88081,
			BlockTimestamp: primitives.TimestampNano(t2.UnixNano()),
		},
		TransactionReceipt: &protocol.TransactionReceiptBuilder{
			Txhash:              h1,
			ExecutionResult:     protocol.EXECUTION_RESULT_ERROR_UNEXPECTED,
			OutputArgumentArray: a3,
		},
		TransactionStatus: protocol.TRANSACTION_STATUS_NO_RECORD_FOUND,
		PackedProof:       []byte{0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99},
	}).Build()
	fmt.Printf(`"GetTransactionReceiptProofResponse": "%s"`+"\n\n", base64.StdEncoding.EncodeToString(r4.Raw()))
}
