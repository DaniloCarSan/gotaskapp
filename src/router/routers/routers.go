package routers

import (
	"net/http"

	"github.com/gorilla/mux"
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

func goThroughRoutes(rgs []RouteGroup, f func(r Route)) {
	for _, rg := range rgs {
		for _, route := range rg.Routes {
			route.URI = rg.Name + route.URI
			f(route)
		}
	}
}

// Publish routers in mux
func PublishInMux(r *mux.Router) *mux.Router {

	goThroughRoutes(
		[]RouteGroup{
			routersAuth,
			routersUser,
			routersTask,
		},
		func(route Route) {
			r.HandleFunc(route.URI, route.Execute).Methods(route.Method)
		},
	)

	return r
}
