package main

import (
	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "example_api_gateway")

	gprcClients, _ := services.NewServicesRepo(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)
}
