package router

import (
	"gotaskapp/src/router/routers"

	"github.com/gorilla/mux"
)

// Publish routers in mux
func PublishInMux(r *mux.Router) *mux.Router {

	routers.GoThroughRoutes(
		[]routers.RouteGroup{
			routers.Auth,
			routers.User,
			routers.Task,
		},
		func(route routers.Route) {
			r.HandleFunc(route.URI, route.Execute).Methods(route.Method)
		},
	)

	return r
}

// Load all routes of the application
func LoadRouters() *mux.Router {
	r := mux.NewRouter()

	return PublishInMux(r)
}
