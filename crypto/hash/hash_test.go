// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package hash

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var someData = []byte("testing")

const (
	ExpectedSha256 = "cf80cd8aed482d5d1527d7dc72fceff84e6326592848447d2dc0b0e87dfc9a90"
)

func TestCalcSha256(t *testing.T) {
	h := CalcSha256(someData)
	require.Equal(t, SHA256_HASH_SIZE_BYTES, len(h))
	require.Equal(t, ExpectedSha256, h.String(), "result should match")
}

func TestCalcSha256_MultipleChunks(t *testing.T) {
	h := CalcSha256(someData[:3], someData[3:])
	require.Equal(t, SHA256_HASH_SIZE_BYTES, len(h))
	require.Equal(t, ExpectedSha256, h.String(), "result should match")
}

func BenchmarkCalcSha256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcSha256(someData)
	}
}
