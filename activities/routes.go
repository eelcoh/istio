package main

import (
	"github.com/eelcoh/go-api"
)

var routes = api.Routes{

	api.NewRoute(
		"Create Activity",
		"POST",
		"/",
		createActivity,
	),

	api.NewRoute(
		"List Activities",
		"GET",
		"/",
		listActivities,
	),

	api.NewRoute(
		"Delete Activity By Refererence",
		"DELETE",
		"/ref/{ref}",
		deleteActivityByReference,
	),

	api.NewRoute(
		"Delete Activity",
		"DELETE",
		"/{id}",
		deleteActivity,
	),

	api.NewRoute(
		"Health",
		"GET",
		"/health",
		health,
	),
}
