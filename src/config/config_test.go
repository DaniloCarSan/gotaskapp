package config

import (
	"os"
	"testing"
)

func TestConfigVariables(t *testing.T) {

	err := Load("./../../.env_exemple")

	if err != nil {
		t.Errorf("Error load variables: %v ", err)
	}

	if _, exists := os.LookupEnv("APP_HOST"); !exists {
		t.Errorf("error variable APP_HOST does not exist in .env_example check the .env file too")
	}

	if _, exists := os.LookupEnv("APP_PORT"); !exists {
		t.Errorf("error variable APP_PORT does not exist in .env_example check the .env file too")
	}
}
