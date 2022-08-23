package main

import (
	"gotaskapp/src/config"
	"gotaskapp/src/router"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load environment variables from .env file
	if err := config.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// Configure sentry client credentials
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SENTRY_DNS,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	app := gin.Default()

	//Inject sentrygin HandlerFunc in gin to make it available in routes
	app.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	router.LoadRouters(app)

	app.Run(config.APP_HOST_FULL)
}
