package main

import (
	"fmt"

	"github.com/rs/xid"
)

// Message is the basic objct.
type Activity struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Activity string `json:"activity"`
	Ref      string `json:"ref"`
}

// NewMessage is the struct used for creating new messages, which do not have an id yet
type NewActivity struct {
	Type     string `json:"type"`
	Activity string `json:"activity"`
	Ref      string `json:"ref"`
}

func add(nwAct NewActivity) {
	id := xid.New().String()
	act := Activity{id, nwAct.Type, nwAct.Activity, nwAct.Ref}
	activities = append(activities, act)
}

func del(id string) {

	for i, a := range activities {
		if a.ID == id {
			activities = append(activities[:i], activities[i+1:]...)
		}
	}
	fmt.Println(activities)

}

func delByRef(ref string) {

	for i, a := range activities {
		if a.Ref == ref {
			activities = append(activities[:i], activities[i+1:]...)
		}
	}
	fmt.Println(activities)

}
