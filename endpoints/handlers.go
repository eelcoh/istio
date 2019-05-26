package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	endpointdiscovery "github.com/eelcoh/endpoint-discovery"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up!\n")
}

// new ..
func registerAPIEndpoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hostname := params["hostname"]
	port, err := strconv.Atoi(params["port"])

	if err != nil {
		http.Error(w, "Unknown port", 400)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var manifest endpointdiscovery.Manifest
	err = json.NewDecoder(r.Body).Decode(&manifest)
	if err != nil {
		http.Error(w, "cannot parse manifest", 400)
		return
	}

	add(manifest, hostname, port)

	w.WriteHeader(201)

}

func findEndpoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	api := params["api"]
	occurence := params["occurence"]
	version := params["version"]

	res := find(api, occurence, version)

	returnEndpointInstances(res, w, r)

}

func returnEndpointInstances(res []endpointdiscovery.EndpointInstance, w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(res)

	if err != nil {

		log.Printf("endpoint instances : %s", res)
		panic(err)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

}

func returnBadRequest(err error, w http.ResponseWriter, r *http.Request) {

	panic(err)

}
