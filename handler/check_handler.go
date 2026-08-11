package handler

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
	SendResponse(c, nil, gin.H{"status": "ok"})
}
