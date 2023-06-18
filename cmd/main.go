package main

import (
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/config"
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/routes"
)

func main() {
	// Load the configuration
	cfg := config.LoadConfig()

	// Setup routes and start the server
	r := routes.SetupRouter(cfg.Server.Port)
	r.Run()
}
