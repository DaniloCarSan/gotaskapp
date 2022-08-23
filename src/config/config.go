package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APP_PORT      = 0
	APP_HOST      = ""
	APP_HOST_FULL = ""
)

// Load application settings
func Load() {

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	APP_PORT, err = strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		APP_PORT = 5000
	}

	APP_HOST = os.Getenv("APP_HOST")

	APP_HOST_FULL = fmt.Sprintf("%s:%d", APP_HOST, APP_PORT)
}
