package router

import (
	"net/http"

	"github.com/bg-vc/vc-gin-api/handler"
	"github.com/bg-vc/vc-gin-api/router/middleware"
	"github.com/gin-gonic/gin"
)

func Load(
	engine *gin.Engine,
	accounts *handler.AccountHandler,
	users *handler.UserHandler,
	auth gin.HandlerFunc,
	logging gin.HandlerFunc,
) *gin.Engine {
	engine.Use(gin.Recovery(), middleware.Headers, logging)
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, handler.Response{Code: 8, Msg: "route does not exist"})
	})

	engine.GET("/api/check/health", handler.HealthCheck)

	account := engine.Group("/api/account")
	account.POST("/login", accounts.Login)
	account.POST("/updatePass", auth, accounts.UpdatePassword)

	engine.GET("/api/user/common/queryAll", users.List)

	protectedUsers := engine.Group("/api/user/auth", auth)
	protectedUsers.POST("/add", users.Create)
	protectedUsers.POST("/update", users.Update)

	return engine
}
