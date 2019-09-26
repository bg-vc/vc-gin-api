package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vc-gin-api/handler"
	"vc-gin-api/router/middleware"
)

const (
	IsAdmin = 1
	IsUser  = 2
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "the incorrect api route")
	})

	check := g.Group("/api/check")
	{
		check.GET("/health", handler.HealthCheck)
		check.GET("/disk", handler.DiskCheck)
		check.GET("/cpu", handler.CPUCheck)
		check.GET("/ram", handler.RAMCheck)
	}

	account := g.Group("/api/operate/account")
	{
		account.POST("/login", handler.LoginAccount)
		account.POST("/updatePass", handler.UpdateAccountPwd)
	}

	user := g.Group("/api/user/common")
	{
		user.GET("/queryAll", handler.QueryUsers)
	}

	userAuth := g.Group("/api/user/auth")
	userAuth.Use(middleware.AuthMiddleware(IsUser))
	{
		userAuth.POST("/add", handler.AddUser)
		userAuth.POST("/update", handler.UpdateUser)
	}

	return g
}
