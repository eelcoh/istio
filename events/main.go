package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/eelcoh/go-api"
)

var events []Event
var port int

func init() {
	fmt.Println("initialising")
	events = make([]Event, 0)
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
