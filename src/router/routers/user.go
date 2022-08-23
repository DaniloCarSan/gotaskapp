package routers

import (
	"net/http"
)

var routersUser = RouteGroup{
	Name: "/users",
	Routes: []Route{
		{
			URI:                   "",
			Method:                http.MethodGet,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
		{
			URI:                   "",
			Method:                http.MethodPatch,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
	},
}
