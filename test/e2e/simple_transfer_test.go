// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package e2e

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleTransfer(t *testing.T) {
	h := newHarness()
	defer h.shutdown()

	// create sender account
	sender, err := orbs.CreateAccount()
	require.NoError(t, err)

	// create receiver account
	receiver, err := orbs.CreateAccount()
	require.NoError(t, err)

	// create client
	endpoint := getEndpoint()
	client := orbs.NewClient(endpoint, VIRTUAL_CHAIN_ID, codec.NETWORK_TYPE_TEST_NET)

	// create transfer transaction
	tx, txId, err := client.CreateTransaction(
		sender.PublicKey,
		sender.PrivateKey,
		"BenchmarkToken",
		"transfer",
		uint64(10), receiver.AddressAsBytes())
	require.NoError(t, err)

	// send the transaction
	transferResponse, err := client.SendTransaction(tx)
	t.Log("Transfer response:")
	t.Logf("%+v", transferResponse)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, transferResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, transferResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, transferResponse.TransactionStatus)

	// check the transaction status
	statusResponse, err := client.GetTransactionStatus(txId)
	t.Log("Status response:")
	t.Logf("%+v", statusResponse)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, statusResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, statusResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, statusResponse.TransactionStatus)

	// check the transaction status receipt proof
	txProofResponse, err := client.GetTransactionReceiptProof(txId)
	t.Log("Receipt proof response:")
	t.Logf("%+v", txProofResponse)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, txProofResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, txProofResponse.ExecutionResult)
	require.Equal(t, codec.TRANSACTION_STATUS_COMMITTED, txProofResponse.TransactionStatus)
	require.True(t, len(txProofResponse.PackedProof) > 20)
	require.True(t, len(txProofResponse.PackedReceipt) > 10)

	// create balance query
	query, err := client.CreateQuery(
		receiver.PublicKey,
		"BenchmarkToken",
		"getBalance",
		receiver.AddressAsBytes())
	require.NoError(t, err)

	// send the query
	balanceResponse, err := client.SendQuery(query)
	t.Log("Query response:")
	t.Logf("%+v", balanceResponse)
	require.NoError(t, err)
	require.Equal(t, codec.REQUEST_STATUS_COMPLETED, balanceResponse.RequestStatus)
	require.Equal(t, codec.EXECUTION_RESULT_SUCCESS, balanceResponse.ExecutionResult)
	require.Equal(t, uint64(10), balanceResponse.OutputArguments[0])

	// get the block which contains the transfer transaction
	blockResponse, err := client.GetBlock(transferResponse.BlockHeight)
	require.NoError(t, err)
	t.Log("Block response:")
	t.Logf("%+v", blockResponse)
	require.Equal(t, transferResponse.BlockHeight, blockResponse.BlockHeight)
	t.Log("  TransactionBlockHeader:")
	t.Logf("  %+v", blockResponse.TransactionsBlockHeader)
	require.Equal(t, transferResponse.BlockHeight, blockResponse.TransactionsBlockHeader.BlockHeight)
	require.EqualValues(t, 1, blockResponse.TransactionsBlockHeader.NumTransactions)
	t.Log("  ResultsBlockHeader:")
	t.Logf("  %+v", blockResponse.ResultsBlockHeader)
	require.Equal(t, transferResponse.BlockHeight, blockResponse.ResultsBlockHeader.BlockHeight)
	require.Equal(t, blockResponse.TransactionsBlockHash, blockResponse.ResultsBlockHeader.TransactionsBlockHash)
	require.EqualValues(t, 1, blockResponse.ResultsBlockHeader.NumTransactionReceipts)
	t.Log("  Transactions:")
	t.Logf("  %+v", blockResponse.Transactions[0])
	require.Equal(t, "BenchmarkToken", blockResponse.Transactions[0].ContractName)
	require.Equal(t, "transfer", blockResponse.Transactions[0].MethodName)
	require.Equal(t, uint64(10), blockResponse.Transactions[0].InputArguments[0])
	require.Equal(t, receiver.AddressAsBytes(), blockResponse.Transactions[0].InputArguments[1])
}
