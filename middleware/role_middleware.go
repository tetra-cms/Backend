package middleware

import (
	"net/http"

	"tetra-server/models"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...models.Role) gin.HandlerFunc {

	return func(c *gin.Context) {

		role, ok := c.Get("role")

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		userRole := role.(models.Role)

		for _, r := range roles {

			if userRole == r {
				c.Next()
				return
			}

		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "forbidden",
		})
	}
}
