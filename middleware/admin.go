package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userRole")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Hanya admin yang boleh mengakses"})
			c.Abort()
			return
		}
		c.Next()
	}
}

