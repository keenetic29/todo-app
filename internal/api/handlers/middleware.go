package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"strings"
	"todo-app/pkg/jwt"
	"todo-app/pkg/logger"
)

func AuthMiddleware(jwtUtil jwt.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Error.Println("Authorization header missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := authHeader
		// if tokenString == authHeader {
		// 	logger.Error.Println("Bearer token not found")
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
		// 	c.Abort()
		// 	return
		// }

		userID, err := jwtUtil.ValidateToken(tokenString)
		if err != nil {
			logger.Error.Printf("Invalid token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		logger.Info.Printf("Authenticated user ID: %d", userID)
		c.Set("userID", userID)
		c.Next()
	}
}