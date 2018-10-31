package codec

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/keys"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/pkg/errors"
	"time"
)

type CallMethodRequest struct {
	ProtocolVersion uint32
	VirtualChainId  uint32
	Timestamp       time.Time
	NetworkType     NetworkType
	PublicKey       []byte
	ContractName    string
	MethodName      string
	InputArguments  []interface{}
}

type CallMethodResponse struct {
	RequestStatus   RequestStatus
	ExecutionResult ExecutionResult
	OutputArguments []interface{}
	BlockHeight     uint64
	BlockTimestamp  time.Time
}

func EncodeCallMethodRequest(req *CallMethodRequest) ([]byte, error) {
	// validate
	if req.ProtocolVersion != 1 {
		return nil, errors.Errorf("expected ProtocolVersion 1, %d given", req.ProtocolVersion)
	}
	if len(req.PublicKey) != keys.ED25519_PUBLIC_KEY_SIZE_BYTES {
		return nil, errors.Errorf("expected PublicKey length %d, %d given", keys.ED25519_PUBLIC_KEY_SIZE_BYTES, len(req.PublicKey))
	}

	// encode method arguments
	inputArgumentArray, err := MethodArgumentsOpaqueEncode(req.InputArguments)
	if err != nil {
		return nil, err
	}

	// encode network type
	networkType, err := networkTypeEncode(req.NetworkType)
	if err != nil {
		return nil, err
	}

	// encode request
	res := (&client.CallMethodRequestBuilder{
		Transaction: &protocol.TransactionBuilder{
			ProtocolVersion: primitives.ProtocolVersion(req.ProtocolVersion),
			VirtualChainId:  primitives.VirtualChainId(req.VirtualChainId),
			Timestamp:       primitives.TimestampNano(req.Timestamp.UnixNano()),
			Signer: &protocol.SignerBuilder{
				Scheme: protocol.SIGNER_SCHEME_EDDSA,
				Eddsa: &protocol.EdDSA01SignerBuilder{
					NetworkType:     networkType,
					SignerPublicKey: primitives.Ed25519PublicKey(req.PublicKey),
				},
			},
			ContractName:       primitives.ContractName(req.ContractName),
			MethodName:         primitives.MethodName(req.MethodName),
			InputArgumentArray: inputArgumentArray,
		},
	}).Build()

	// return
	return res.Raw(), nil
}

func DecodeCallMethodResponse(buf []byte) (*CallMethodResponse, error) {
	// decode response
	res := client.CallMethodResponseReader(buf)
	if !res.IsValid() {
		return nil, errors.New("response is corrupt and cannot be decoded")
	}

	// decode request status
	requestStatus, err := requestStatusDecode(res.RequestStatus())
	if err != nil {
		return nil, err
	}

	// decode execution result
	executionResult, err := executionResultDecode(res.CallMethodResult())
	if err != nil {
		return nil, err
	}

	// decode method arguments
	outputArgumentArray, err := MethodArgumentsOpaqueDecode(res.RawOutputArgumentArrayWithHeader())
	if err != nil {
		return nil, err
	}

	// return
	return &CallMethodResponse{
		RequestStatus:   requestStatus,
		ExecutionResult: executionResult,
		OutputArguments: outputArgumentArray,
		BlockHeight:     uint64(res.BlockHeight()),
		BlockTimestamp:  time.Unix(0, int64(res.BlockTimestamp())),
	}, nil
}
