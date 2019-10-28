package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vc-gin-api/pkg/errno"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
