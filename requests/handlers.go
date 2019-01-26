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
func createRequest(w http.ResponseWriter, r *http.Request) {
	var reqBody NewRequest

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println("error")

		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(reqBody)

	req := add(reqBody)
	logNewRequest(req)

	returnRequests(w, r)

}

func listRequests(w http.ResponseWriter, r *http.Request) {

	returnRequests(w, r)

}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Print("deleting ..")
	params := mux.Vars(r)
	fmt.Println(params)
	id := params["id"]

	del(id)

	returnRequests(w, r)

}

func returnRequests(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(requests)

	if err != nil {

		log.Printf("service-requests : %s", requests)
		panic(err)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

}

func logNewRequest(req Request) {

	t := "Service Request"
	actStr := fmt.Sprintf("[Create][%s] by %s : %s", req.Category, req.Name, req.Request)
	ref := fmt.Sprintf("/service-requests/%s", req.ID)

	logActivity(t, actStr, ref)
}

func logDeleteRequest(id string) {
	t := "Service Request"
	actStr := fmt.Sprintf("[Delete][%s]", id)
	ref := fmt.Sprintf("/service-requests/%s", id)

	logActivity(t, actStr, ref)
}
