package main

import (
	"github.com/eelcoh/go-api"
)

var routes = api.Routes{

	api.NewRoute(
		"Create Token",
		"POST",
		"/tokens",
		createToken,
	),

	api.NewRoute(
		"Health",
		"GET",
		"/health",
		health,
	),
}
