package main

import (
	"github.com/eelcoh/go-api"
)

var routes = api.Routes{

	api.NewRoute(
		"Create Event",
		"POST",
		"/",
		createEvent,
	),

	api.NewRoute(
		"List Events",
		"GET",
		"/",
		listEvents,
	),

	api.NewRoute(
		"Delete Event By Refererence",
		"DELETE",
		"/ref/{ref}",
		deleteEventByReference,
	),

	api.NewRoute(
		"Delete Event",
		"DELETE",
		"/{id}",
		deleteEvent,
	),

	api.NewRoute(
		"Health",
		"GET",
		"/health",
		health,
	),
}
