package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/controllers"
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/middleware"
)

// SetupRouter sets up the Gin router with routes
func SetupRouter(port string) *gin.Engine {
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
