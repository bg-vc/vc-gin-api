package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vc-gin-api/handler"
	"vc-gin-api/router/middleware"
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

	demo := g.Group("/api/demo")
	{
		demo.GET("/query", handler.QueryDemo)
	}

	return g
}