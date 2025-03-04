package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware verifies the JWT token and extracts user claims.
func JWTAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token provided"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
			c.Abort()
			return
		}

		// Extract claims (Assume payload contains `id`, `role`, `tenant_id`)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid claims"})
			c.Abort()
			return
		}

		userID, role, tenantID := claims["id"], claims["role"], claims["tenant_id"]
		if userID == nil || role == nil || tenantID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Missing claims"})
			c.Abort()
			return
		}

		// Store claims in context
		c.Set("user_id", userID)
		c.Set("role", role)
		c.Set("tenant_id", tenantID)

		c.Next()
	}
}
