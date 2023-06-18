package routes

import (
	"Karthika-Rajagopal/go-restful-backend/internal/controllers"
	"Karthika-Rajagopal/go-restful-backend/internal/middleware"
	"Karthika-Rajagopal/go-restful-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router and routes for the application
func SetupRouter(userRepo repositories.UserRepository) *gin.Engine {
	r := gin.Default()

	authController := controllers.NewAuthController(userRepo)
	userController := controllers.NewUserController(userRepo)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/profile", userController.GetProfile)
		userRoutes.PUT("/profile", userController.UpdateProfile)
	}

	return r
}
