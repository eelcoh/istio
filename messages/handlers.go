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
func createMessage(w http.ResponseWriter, r *http.Request) {
	var msgBody NewMessage

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&msgBody)
	if err != nil {
		fmt.Println("error")

		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(msgBody)

	msg := add(msgBody)
	logNew(msg)

	returnMessages(w, r)

}

func listMessages(w http.ResponseWriter, r *http.Request) {

	returnMessages(w, r)

}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("deleting ..")
	params := mux.Vars(r)
	fmt.Println(params)
	id := params["id"]

	erase := r.URL.Query().Get("erase")

	del(id)

	if erase == "true" {
		deleteLog(id)
	} else {
		logDelete(id)
	}

	returnMessages(w, r)

}

func returnMessages(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(messages)

	if err != nil {

		log.Printf("messages : %s", messages)
		panic(err)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

}

func logNew(msg Message) {

	t := "Message"
	actStr := fmt.Sprintf("[Create] by %s : %s", msg.Name, msg.Message)
	ref := fmt.Sprintf("%s", msg.ID)

	logActivity(t, actStr, ref)
}

func logDelete(id string) {
	t := "Message"
	actStr := fmt.Sprintf("[Delete]", id)
	ref := fmt.Sprintf("%s", id)

	logActivity(t, actStr, ref)
}
