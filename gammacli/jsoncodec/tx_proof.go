package jsoncodec

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"strconv"
)

func MarshalTxProofResponse(r *codec.GetTransactionReceiptProofResponse) ([]byte, error) {
	return json.MarshalIndent(&struct {
		RequestStatus     codec.RequestStatus
		PackedProof       string
		TransactionStatus codec.TransactionStatus
		BlockHeight       string
		BlockTimestamp    string
	}{
		RequestStatus:     r.RequestStatus,
		PackedProof:       hex.EncodeToString(r.PackedProof),
		TransactionStatus: r.TransactionStatus,
		BlockHeight:       strconv.FormatUint(r.BlockHeight, 10),
		BlockTimestamp:    r.BlockTimestamp.UTC().Format(codec.ISO_DATE_FORMAT),
	}, "", "  ")
}
