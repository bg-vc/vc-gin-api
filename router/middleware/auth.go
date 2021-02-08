package middleware

import (
	"github.com/gin-gonic/gin"
	"vc-gin-api/handler"
	"vc-gin-api/pkg"
	"vc-gin-api/pkg/errno"
	"vc-gin-api/pkg/log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := pkg.ParseRequest(c); err != nil {
			log.Errorf(err, "pkg.ParseRequest error")
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
