package main

import (
	"github.com/rs/xid"
)

// Message is the basic objct.
type Message struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// NewMessage is the struct used for creating new messages, which do not have an id yet
type NewMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func add(nwMsg NewMessage) Message {
	id := xid.New().String()
	msg := Message{id, nwMsg.Name, nwMsg.Message}
	messages[id] = msg
	return msg
}

func del(id string) {
	delete(messages, id)
}
