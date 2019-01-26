package main

import (
	"fmt"

	"github.com/rs/xid"
)

// Request is the basic object.
type Request struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Request  string `json:"request"`
}

// NewRequest is the struct used for creating new service requests, which do not have an id yet
type NewRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Request  string `json:"request"`
}

func add(nwSR NewRequest) Request {
	id := xid.New().String()
	req := Request{id, nwSR.Name, nwSR.Category, nwSR.Request}
	requests = append(requests, req)

	return req
}

func del(id string) {

	for i, a := range requests {
		if a.ID == id {
			requests = append(requests[:i], requests[i+1:]...)
		}
	}
	fmt.Println(requests)

}
