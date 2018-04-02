package models

import (
	"container/list"
)

type EventType int

// event code
const (
	EVENT_JOIN    = 0
	EVENT_LEAVE   = 1
	EVENT_MESSAGE = 2
)

type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	User      string
	Timestamp int // event time
	Content   string
}

const eventStoreSize = 20

var eventLists = list.New()

// AddEvent saves new event to events list.
func AddEvent(event Event) {
	if eventLists.Len() >= eventStoreSize {
		eventLists.Remove(eventLists.Front())
	}
	eventLists.PushBack(event)
}

//历史记录
// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, eventLists.Len())
	for event := eventLists.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
