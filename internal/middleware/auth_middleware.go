package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Karthika-Rajagopal/go-restful-backend/internal/utils"
)

// AuthMiddleware is a middleware function for JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		userID, err := utils.ValidateJWTToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
