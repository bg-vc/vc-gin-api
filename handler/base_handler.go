package handler

import (
	"log/slog"

	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SendResponse(c *gin.Context, err error, data any) {
	status, code, message := errno.Decode(err)
	if status >= 500 && err != nil {
		slog.Error("request failed", "error", err)
	}
	c.JSON(status, Response{Code: code, Msg: message, Data: data})
}
