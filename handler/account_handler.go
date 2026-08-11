package handler

import (
	"github.com/bg-vc/vc-gin-api/model"
	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"github.com/bg-vc/vc-gin-api/router/middleware"
	"github.com/bg-vc/vc-gin-api/service"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accounts *service.AccountService
}

func NewAccountHandler(accounts *service.AccountService) *AccountHandler {
	return &AccountHandler{accounts: accounts}
}

func (handler *AccountHandler) Login(c *gin.Context) {
	var form model.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	response, err := handler.accounts.Login(c.Request.Context(), form)
	SendResponse(c, err, response)
}

func (handler *AccountHandler) UpdatePassword(c *gin.Context) {
	var form model.UpdatePasswordForm
	if err := c.ShouldBindJSON(&form); err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	username := middleware.Username(c)
	err := handler.accounts.UpdatePassword(c.Request.Context(), username, form)
	SendResponse(c, err, nil)
}
