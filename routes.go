package main

import "net/http"

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
		"/accounts",
		AccountIndex,
	},
	Route{
		"AccountId",
		"GET",
		"/accounts/{accountId}",
		AccountId,
	},
	Route{
		"AccountCreate",
		"POST",
		"/accounts",
		AccountCreate,
	},
}
