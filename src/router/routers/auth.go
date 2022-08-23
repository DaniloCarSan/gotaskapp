package routers

import (
	"net/http"
)

var Auth = RouteGroup{
	Name: "/auth",
	Routes: []Route{
		{
			URI:                   "/sign/in",
			Method:                http.MethodPost,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: false,
		},
		{
			URI:                   "/sign/up",
			Method:                http.MethodPost,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: false,
		},
		{
			URI:                   "/password/reset",
			Method:                http.MethodPost,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: false,
		},
		{
			URI:                   "/token/renew",
			Method:                http.MethodGet,
			Execute:               func(http.ResponseWriter, *http.Request) {},
			RequireAuthentication: true,
		},
	},
}
