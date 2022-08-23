package router

import (
	"gotaskapp/src/router/routers"

	"github.com/gorilla/mux"
)

// Load all routes of the application
func LoadRouters() *mux.Router {
	r := mux.NewRouter()

	return routers.PublishInMux(r)
}
