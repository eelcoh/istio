package main

import (
	"fmt"

	"github.com/rs/xid"
)

// Message is the basic objct.
type Event struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Event 	 string `json:"Event"`
	Ref      string `json:"ref"`
}

// NewMessage is the struct used for creating new messages, which do not have an id yet
type NewEvent struct {
	Type     string `json:"type"`
	Event 	 string `json:"Event"`
	Ref      string `json:"ref"`
}

func add(nwEvt NewEvent) {
	id := xid.New().String()
	Evt := Event{id, nwEvt.Type, nwEvt.Event, nwEvt.Ref}
	events = append(events, Evt)
}

func del(id string) {

	for i, a := range events {
		if a.ID == id {
			events = append(events[:i], events[i+1:]...)
		}
	}
	fmt.Println(events)

}

func delByRef(ref string) {

	for i, a := range events {
		if a.Ref == ref {
			events = append(events[:i], events[i+1:]...)
		}
	}
	fmt.Println(events)

}
