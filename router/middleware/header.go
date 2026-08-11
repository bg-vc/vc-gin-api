package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Headers(c *gin.Context) {
	c.Header("Cache-Control", "no-store")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-Frame-Options", "DENY")
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
