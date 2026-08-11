package middleware

import (
	"github.com/bg-vc/vc-gin-api/pkg"
	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"github.com/gin-gonic/gin"
)

const usernameKey = "authenticatedUsername"

func Auth(tokens *pkg.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenContext, err := tokens.ParseAuthorization(c.GetHeader("Authorization"))
		if err != nil {
			status, code, message := errno.Decode(errno.ErrTokenInvalid)
			c.AbortWithStatusJSON(status, gin.H{"code": code, "msg": message, "data": nil})
			return
		}
		c.Set(usernameKey, tokenContext.Username)
		c.Next()
	}
}

func Username(c *gin.Context) string {
	username, _ := c.Get(usernameKey)
	value, _ := username.(string)
	return value
}
