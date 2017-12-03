package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AccountIndex",
		"GET",
		"/Accounts",
		AccountIndex,
	},
	Route{
		"AccountId",
		"GET",
		"/Accounts/{AccountId}",
		AccountId,
	},
}
