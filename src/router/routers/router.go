package routers

import (
	"net/http"
)

// Represents Route object
type Route struct {
	URI                   string
	Method                string
	Execute               func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// Represents a route group object that contais routes children
type RouteGroup struct {
	Name   string
	Routes []Route
}

// Go through routes
func GoThroughRoutes(rgs []RouteGroup, f func(r Route)) {
	for _, rg := range rgs {
		for _, route := range rg.Routes {
			route.URI = rg.Name + route.URI
			f(route)
		}
	}
}
