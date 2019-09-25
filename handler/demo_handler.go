package handler

import (
	"github.com/gin-gonic/gin"
	"vc-gin-api/pkg"
)

func QueryDemo(c *gin.Context) {
	demo := &Demo{
		Name:    "Vince",
		Age:     28,
		Address: "beijing",
	}
	pkg.Log.Infof("demo1:#%v", demo)
	pkg.Log.Infof("demo2:#%v", demo)
	SendResponse(c, nil, demo)

}

type Demo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
