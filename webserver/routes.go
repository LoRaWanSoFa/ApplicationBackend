package webserver

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
		"TodoIndex",
		"GET",
		"/todos",
		MessageIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		MessageCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		MessageShow,
	},
}
