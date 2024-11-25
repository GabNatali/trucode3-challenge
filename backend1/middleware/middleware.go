package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GabNatali/trucode3-challenge-final/internal/crypto"
	"github.com/gin-gonic/gin"
)

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

		crypto, exists := c.MustGet("crypto").(crypto.Crypto)
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get crypto"})
			return
		}

		secret := c.MustGet("secret").(string)
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

		userId, err := crypto.ParseAndValidateJWT(tokenParts[1], secret)

		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}

		fmt.Println(userId)
		c.Set("userId", userId)
		c.Next()
	}
}
