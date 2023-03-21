package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/SENERGY-Platform/kong_admin/internal/api"
	"github.com/SENERGY-Platform/kong_admin/internal/config"
	"github.com/SENERGY-Platform/kong_admin/internal/kong_client"
)

var endpoints = map[string]func(kongClient *kong_client.KongClient, configuration *config.Config) http.Handler{
	"/routes": func(kongClient *kong_client.KongClient, configuration *config.Config) http.Handler {
		return api.New(kongClient, configuration)
	},
}

func registerHandlers(kongClient *kong_client.KongClient, configuration *config.Config) error {
	for path, getHandler := range endpoints {
		log.Printf("Register endpoint %s", path)
		handler := getHandler(kongClient, configuration)
		http.Handle(path, handler)
	}
	return nil
}

func StartServer(kongClient *kong_client.KongClient, configuration *config.Config) {
	err := registerHandlers(kongClient, configuration)
	if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	log.Printf("Start Server")
	err = http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config")
	}
	kongClient := kong_client.LoadKongClient(config)
	StartServer(kongClient, config)
}
