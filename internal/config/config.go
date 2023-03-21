package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	KONG_ADMIN_URL string
}

func LoadConfig() (*Config, error) {
	adminURL, ok := os.LookupEnv("KONG_ADMIN_URL")
	if !ok {
		log.Println("Missing environment variable")
		return nil, errors.New("Missing environment variable")
	}
	config := &Config{
		KONG_ADMIN_URL: adminURL,
	}
	return config, nil
}
