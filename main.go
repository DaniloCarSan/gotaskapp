package main

import (
	"gotaskapp/src/config"
	"gotaskapp/src/router"
	"log"
	"net/http"
)

func main() {

	if err := config.Load(".env"); err != nil {
		log.Fatal(err)
	}

	r := router.LoadRouters()

	log.Fatal(http.ListenAndServe(config.APP_HOST_FULL, r))
}
