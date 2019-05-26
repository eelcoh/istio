package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up!\n")
}

// new ..
func createEvent(w http.ResponseWriter, r *http.Request) {
	var eventBody NewEvent

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&eventBody)
	if err != nil {
		fmt.Println("error")

		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(eventBody)

	add(eventBody)

	returnEvents(w, r)

}

func listEvents(w http.ResponseWriter, r *http.Request) {

	returnEvents(w, r)

}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Print("deleting ..")
	params := mux.Vars(r)
	fmt.Println(params)
	id := params["id"]

	del(id)

	returnEvents(w, r)

}

func deleteEventByReference(w http.ResponseWriter, r *http.Request) {
	fmt.Print("deleting by ref..")
	params := mux.Vars(r)
	fmt.Println(params)
	ref := params["ref"]

	delByRef(ref)

	returnEvents(w, r)

}
func returnEvents(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(events)

	if err != nil {

		log.Printf("events : %s", events)
		panic(err)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

}
