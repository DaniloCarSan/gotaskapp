package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_PORT      = ""
	APP_HOST      = ""
	APP_HOST_FULL = ""
	SENTRY_DNS    = ""
)

// Load application settings
func Load(fileEnv string) error {

	var err error
	var value string
	var exists bool

	if err = godotenv.Load(fileEnv); err != nil {
		return err
	}

	value, exists = os.LookupEnv("APP_HOST")

	if !exists {
		return errors.New("not found variable APP_HOST in .env file")
	}

	APP_HOST = value

	value, exists = os.LookupEnv("APP_PORT")

	if !exists {
		return errors.New("not found variable APP_PORT in .env file")
	}

	APP_PORT = value

	APP_HOST_FULL = fmt.Sprintf("%s:%s", APP_HOST, APP_PORT)

	value, exists = os.LookupEnv("SENTRY_DNS")

	if !exists {
		return errors.New("not found variable APP_PORT in .env file")
	}

	if value == "" {
		return errors.New("required a value for variable SENTRY in .env file")
	}

	SENTRY_DNS = value

	return nil
}
