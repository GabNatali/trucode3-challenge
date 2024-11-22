package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GabNatali/trucode3-challenge-final/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Payload struct {
	jwt.MapClaims        // ExpiryAt, IssueAt
	Session       string `json:"session"`
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		authService, exists := c.MustGet("authService").(auth.AuthService)
		if !exists {
			c.JSON(500, gin.H{"message": "error server"})
			c.Abort()
			return
		}

		tokenBearer := c.Request.Header.Get("Authorization")

		if tokenBearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		tokenParts := strings.Split(tokenBearer, "Bearer ")
		if len(tokenParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		fmt.Println(tokenParts[1])
		userId, err := authService.VerifyAccessToken(tokenParts[1])

		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}

		fmt.Println(userId)
		c.Set("userId", userId)
		c.Next()
	}
}
