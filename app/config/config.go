package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APP_PORT       = ""
	APP_HOST       = ""
	APP_HOST_FULL  = ""
	APP_URL        = ""
	SENTRY_DNS     = ""
	DB_DRIVE       = "mysql"
	DB_ADDR        = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	DB_USER        = ""
	DB_PASS        = ""
	DB_HOST        = ""
	DB_PORT        = ""
	DB_NAME        = ""
	JWT_SECRET     = ""
	EMAIL_HOST     = ""
	EMAIL_FROM     = ""
	EMAIL_PASSWORD = ""
	EMAIL_PORT     = 0
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
	if value == "" || !exists {
		return errors.New("not found variable APP_HOST in .env file or variable empty")
	}
	APP_HOST = value

	value, exists = os.LookupEnv("APP_PORT")
	if value == "" || !exists {
		return errors.New("not found variable APP_PORT in .env file or variable empty")
	}
	APP_PORT = value

	APP_HOST_FULL = fmt.Sprintf("%s:%s", APP_HOST, APP_PORT)

	APP_URL = fmt.Sprintf("http://%s", APP_HOST_FULL)

	value, exists = os.LookupEnv("SENTRY_DNS")
	if value == "" || !exists {
		return errors.New("not found variable APP_PORT in .env file or variable empty")
	}
	SENTRY_DNS = value

	value, exists = os.LookupEnv("DB_USER")
	if value == "" || !exists {
		return errors.New("not found variable DB_USER in .env file or variable empty")
	}
	DB_USER = value

	value, exists = os.LookupEnv("DB_PASS")
	if value == "" || !exists {
		return errors.New("not found variable DB_PASS in .env file or variable empty")
	}
	DB_PASS = value

	value, exists = os.LookupEnv("DB_HOST")
	if value == "" || !exists {
		return errors.New("not found variable DB_HOST in .env file or variable empty")
	}
	DB_HOST = value

	value, exists = os.LookupEnv("DB_PORT")
	if value == "" || !exists {
		return errors.New("not found variable DB_PORT in .env file or variable empty")
	}
	DB_PORT = value

	value, exists = os.LookupEnv("DB_NAME")
	if value == "" || !exists {
		return errors.New("not found variable DB_NAME in .env file or variable empty")
	}
	DB_NAME = value

	DB_ADDR = fmt.Sprintf(DB_ADDR,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	value, exists = os.LookupEnv("JWT_SECRET")
	if value == "" || !exists {
		return errors.New("not found variable JWT_SECRET in .env file or variable empty")
	}
	JWT_SECRET = value

	value, exists = os.LookupEnv("EMAIL_HOST")
	if value == "" || !exists {
		return errors.New("not found variable EMAIL_HOST in .env file or variable empty")
	}
	EMAIL_HOST = value

	value, exists = os.LookupEnv("EMAIL_FROM")
	if value == "" || !exists {
		return errors.New("not found variable EMAIL_FROM in .env file or variable empty")
	}
	EMAIL_FROM = value

	value, exists = os.LookupEnv("EMAIL_PASSWORD")
	if value == "" || !exists {
		return errors.New("not found variable EMAIL_PASSWORD in .env file or variable empty")
	}
	EMAIL_PASSWORD = value

	value, exists = os.LookupEnv("EMAIL_PORT")
	if value == "" || !exists {
		return errors.New("not found variable EMAIL_PORT in .env file or variable empty")
	}

	EMAIL_PORT, err = strconv.Atoi(value)
	if err != nil {
		return err
	}

	return nil
}
