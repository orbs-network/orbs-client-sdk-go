package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInc(t *testing.T) {
	unit.InServiceScope(nil, nil, func(mockery unit.Mockery) {
		mockery.MockEmitEvent(Inc, uint64(1))

		inc()
		require.EqualValues(t, 1, value())
	})
}