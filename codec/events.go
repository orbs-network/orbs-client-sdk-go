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
