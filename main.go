package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vc-gin-api/pkg"
	"vc-gin-api/router"
)

func main() {
	pkg.InitLogger()
	gin.SetMode("release")
	g := gin.New()

	router.Load(
		g,
	)
	//go AAA()
	pkg.Log.Info("server start:8080")
	pkg.Log.Infof(":%v", http.ListenAndServe(":8080", g).Error())
}

func AAA() {
	for {
		pkg.Log.Info("hello vc")
	}
}
