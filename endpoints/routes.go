package main

import (
	api "github.com/eelcoh/go-api"
)

var routes = api.Routes{

	api.NewRoute(
		"Register API",
		"POST",
		"/{hostname}/{port}",
		registerAPIEndpoints,
	),

	api.NewRoute(
		"Discover Addresses for Endpoints",
		"GET",
		"/{api}/{occurence}/{version}",
		findEndpoints,
	),

	// api.NewRoute(
	// 	"Delete Endpoint",
	// 	"DELETE",
	// 	"/",
	// 	deleteEndpoint,
	// ),

	api.NewRoute(
		"Health",
		"GET",
		"/health",
		health,
	),
}
