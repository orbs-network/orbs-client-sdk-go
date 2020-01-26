package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/events"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

func inc() uint64 {
	v := value() + 1
	state.WriteUint64(COUNTER_KEY, v)
	events.EmitEvent(Inc, v)
	return v
}

func value() uint64 {
	return state.ReadUint64(COUNTER_KEY)
}

func getOwnAddress() []byte {
	return address.GetOwnAddress()
}