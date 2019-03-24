// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

type Event struct {
	ContractName string
	EventName    string
	Arguments    []interface{}
}

func PackedEventsEncode(eventBuilders []*protocol.EventBuilder) []byte {
	eventsArray := (&protocol.EventsArrayBuilder{Events: eventBuilders}).Build()
	return eventsArray.RawEventsArray()
}

func PackedEventsDecode(buf []byte) (res []*Event, err error) {
	res = []*Event{}
	eventsArray := protocol.EventsArrayReader(buf)
	index := 0
	for i := eventsArray.EventsIterator(); i.HasNext(); {
		event := i.NextEvents()
		args, err := PackedArgumentsDecode(event.RawOutputArgumentArrayWithHeader())
		if err != nil {
			return nil, errors.Wrapf(err, "received event %d has invalid arguments", index)
		}
		outputEvent := &Event{
			ContractName: string(event.ContractName()),
			EventName:    string(event.EventName()),
			Arguments:    args,
		}
		res = append(res, outputEvent)
		index++
	}
	return
}
