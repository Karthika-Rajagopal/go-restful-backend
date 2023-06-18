package main

import (
	"log"

	//"github.com/gin-gonic/gin"
	"Karthika-Rajagopal/go-restful-backend/internal/config"
	"Karthika-Rajagopal/go-restful-backend/internal/repositories"
	"Karthika-Rajagopal/go-restful-backend/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg:= config.LoadConfig()

	dsn := "host=localhost user=postgres password=postgres dbname=golang-gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	userRepo := repositories.NewUserRepository(db)

	r := routes.SetupRouter(userRepo)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
