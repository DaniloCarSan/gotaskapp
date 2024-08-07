package main

import (
	"gotaskapp/app/config"
	"gotaskapp/app/router"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Inject sentrygin HandlerFunc in gin to make it available in routes
	app.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	router.LoadRouters(app)

	app.Run(config.APP_HOST_FULL)
}
