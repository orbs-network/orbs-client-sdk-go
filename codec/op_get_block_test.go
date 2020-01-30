package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/stretchr/testify/require"
	"testing"
)
func TestDecodeGetBlockResponse(t *testing.T) {
	block := (&client.GetBlockResponseBuilder{
		RequestResult: &client.RequestResultBuilder{
			RequestStatus: protocol.REQUEST_STATUS_COMPLETED,
		},
		SignedTransactions: []*protocol.SignedTransactionBuilder{
			{
				Transaction: &protocol.TransactionBuilder{
					ContractName: "Movies",
					MethodName: "release",
					Signer: &protocol.SignerBuilder{
						Scheme: protocol.SIGNER_SCHEME_EDDSA,
						Eddsa: &protocol.EdDSA01SignerBuilder{
							SignerPublicKey: primitives.Ed25519PublicKey([]byte("some-public-key-")),
							NetworkType: protocol.NETWORK_TYPE_MAIN_NET,
						},
					},
				},
			},
		},
	}).Build()

	res, err := DecodeGetBlockResponse(block.Raw())
	require.NoError(t, err)
	require.NotNil(t, res.Transactions[0].SignerPublicKey)
}

func TestDecodeGetBlockResponseWithEmptySigner(t *testing.T) {
	block := (&client.GetBlockResponseBuilder{
		RequestResult: &client.RequestResultBuilder{
			RequestStatus: protocol.REQUEST_STATUS_COMPLETED,
		},
		SignedTransactions: []*protocol.SignedTransactionBuilder{
			{
				Transaction: &protocol.TransactionBuilder{
					ContractName: "Movies",
					MethodName: "release",
			},
		},
	},
	}).Build()

	res, err := DecodeGetBlockResponse(block.Raw())
	require.NoError(t, err)
	require.Nil(t, res.Transactions[0].SignerPublicKey)
}