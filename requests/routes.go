package main

import (
	"github.com/eelcoh/go-api"
)

var routes = api.Routes{

	api.NewRoute(
		"Create Request",
		"POST",
		"/",
		createRequest,
	),

	api.NewRoute(
		"List Requests",
		"GET",
		"/",
		listRequests,
	),

	api.NewRoute(
		"Delete Request",
		"DELETE",
		"/{id}",
		deleteRequest,
	),

	api.NewRoute(
		"Health",
		"GET",
		"/health",
		health,
	),
}
