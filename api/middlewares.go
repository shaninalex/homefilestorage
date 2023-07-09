package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Request.Header.Get("X-User")
		if user_id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("user_id", user_id)
		c.Next()
	}
}
