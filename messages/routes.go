package main

import (
	"github.com/eelcoh/go-api"
)

var routes = api.Routes{

	api.NewRoute(
		"Create Message",
		"POST",
		"/",
		createMessage,
	),

	api.NewRoute(
		"List Messages",
		"GET",
		"/",
		listMessages,
	),

	api.NewRoute(
		"Delete Message",
		"DELETE",
		"/{id}",
		deleteMessage,
	),

	api.NewRoute(
		"Health",
		"GET",
		"/health",
		health,
	),
}
