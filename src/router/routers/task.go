package routers

import (
	"net/http"
)

var Task = RouteGroup{
	Name: "/tasks",
	Routes: []Route{
		{
			URI:                   "",
			Method:                http.MethodPost,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
		{
			URI:                   "",
			Method:                http.MethodGet,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
		{
			URI:                   "/{id:[0-9]+}",
			Method:                http.MethodGet,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
		{
			URI:                   "",
			Method:                http.MethodPut,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
		{
			URI:                   "/{id:[0-9]+}",
			Method:                http.MethodDelete,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
		{
			URI:                   "/toggle/done/{id:[0-9]+}",
			Method:                http.MethodPatch,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
	},
}
