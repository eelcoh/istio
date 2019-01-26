package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	api "github.com/eelcoh/go-api"
)

var passphrase string
var port int

var activitiesUrl string

func init() {
	fmt.Println("initialising")
	flag.IntVar(&port, "port", 8080, "port the application listens to")
	flag.Parse()

	activitiesUrl = os.Getenv("ACTIVITIES_URL") // set for dev
	if len(activitiesUrl) == 0 {
		activitiesUrl = "http://activities/" // on k8s
	}
}

func main() {
	fmt.Println("starting")

	fmt.Println("activities url set to", activitiesUrl)

	router := api.NewRouter(routes)
	fmt.Println("routes defined")

	portStr := fmt.Sprintf(":%d", port)
	fmt.Println("exposing port", port)

	log.Fatal(http.ListenAndServe(portStr, router))
	fmt.Println("running")
}
