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
func createActivity(w http.ResponseWriter, r *http.Request) {
	var actBody NewActivity

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&actBody)
	if err != nil {
		fmt.Println("error")

		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(actBody)

	add(actBody)

	returnActivities(w, r)

}

func listActivities(w http.ResponseWriter, r *http.Request) {

	returnActivities(w, r)

}

func deleteActivity(w http.ResponseWriter, r *http.Request) {
	fmt.Print("deleting ..")
	params := mux.Vars(r)
	fmt.Println(params)
	id := params["id"]

	del(id)

	returnActivities(w, r)

}

func returnActivities(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(activities)

	if err != nil {

		log.Printf("activities : %s", activities)
		panic(err)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

}
