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

	a1, _ := codec.MethodArgumentsOpaqueEncode([]interface{}{uint32(1), uint64(2), "hello", []byte{0x01, 0x02, 0x03}})
	a2, _ := codec.MethodArgumentsOpaqueEncode([]interface{}{})
	a3, _ := codec.MethodArgumentsOpaqueEncode([]interface{}{uint64(math.MaxUint64 - 1000)})

	r1 := (&client.SendTransactionResponseBuilder{
		RequestStatus: protocol.REQUEST_STATUS_COMPLETED,
		TransactionReceipt: &protocol.TransactionReceiptBuilder{
			Txhash:              h1,
			ExecutionResult:     protocol.EXECUTION_RESULT_ERROR_SMART_CONTRACT,
			OutputArgumentArray: a1,
		},
		TransactionStatus: protocol.TRANSACTION_STATUS_COMMITTED,
		BlockHeight:       2135,
		BlockTimestamp:    primitives.TimestampNano(t1.UnixNano()),
	}).Build()
	fmt.Printf("%s\n\n", base64.StdEncoding.EncodeToString(r1.Raw()))

	r2 := (&client.CallMethodResponseBuilder{
		RequestStatus:       protocol.REQUEST_STATUS_NOT_FOUND,
		OutputArgumentArray: a2,
		CallMethodResult:    protocol.EXECUTION_RESULT_SUCCESS,
		BlockHeight:         87438,
		BlockTimestamp:      primitives.TimestampNano(t1.UnixNano()),
	}).Build()
	fmt.Printf("%s\n\n", base64.StdEncoding.EncodeToString(r2.Raw()))

	r3 := (&client.GetTransactionStatusResponseBuilder{
		RequestStatus: protocol.REQUEST_STATUS_REJECTED,
		TransactionReceipt: &protocol.TransactionReceiptBuilder{
			Txhash:              h1,
			ExecutionResult:     protocol.EXECUTION_RESULT_ERROR_INPUT,
			OutputArgumentArray: a3,
		},
		TransactionStatus: protocol.TRANSACTION_STATUS_REJECTED_UNKNOWN_SIGNER_SCHEME,
		BlockHeight:       math.MaxUint64 - 5000,
		BlockTimestamp:    primitives.TimestampNano(t2.UnixNano()),
	}).Build()
	fmt.Printf("%s\n\n", base64.StdEncoding.EncodeToString(r3.Raw()))
}