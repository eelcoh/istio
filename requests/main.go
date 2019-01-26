package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eelcoh/go-api"
)

var requests []Request
var port int

var activitiesUrl string

func init() {
	fmt.Println("initialising")
	requests = make([]Request, 0)
	flag.IntVar(&port, "port", 8080, "port the application listens to")
	flag.Parse()

	activitiesUrl = os.Getenv("ACTIVITIES_URL")
	if len(activitiesUrl) == 0 {
		activitiesUrl = "http://activities/"
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
