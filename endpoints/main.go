package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	endpointdiscovery "github.com/eelcoh/endpoint-discovery"
	api "github.com/eelcoh/go-api"
)

var apiEndpoints map[string][]endpointdiscovery.EndpointInstance
var port int

func init() {
	fmt.Println("initialising")
	apiEndpoints = make(map[string][]endpointdiscovery.EndpointInstance)
	flag.IntVar(&port, "port", 8080, "port the application listens to")
	flag.Parse()
}

func main() {
	fmt.Println("starting")

	router := api.NewRouter(routes)
	fmt.Println("routes defined")

	portStr := fmt.Sprintf(":%d", port)
	fmt.Println("exposing port", port)

	log.Fatal(http.ListenAndServe(portStr, router))
	fmt.Println("running")
}
