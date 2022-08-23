package main

import (
	"gotaskapp/src/config"
	"gotaskapp/src/router"
	"log"
	"net/http"
)

func main() {

	config.Load()

	r := router.LoadRouters()

	log.Fatal(http.ListenAndServe(config.APP_HOST_FULL, r))
}
