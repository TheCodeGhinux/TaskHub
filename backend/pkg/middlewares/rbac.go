package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RBAC ensures that only users with allowed roles can access the endpoint.
func RBAC(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Role not found"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next() // Proceed with the request
				return
			}
		}

		// Deny access
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - Insufficient permissions"})
		c.Abort()
	}
}
