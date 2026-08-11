package handler

import (
	"github.com/bg-vc/vc-gin-api/model"
	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"github.com/bg-vc/vc-gin-api/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	users *service.UserService
}

func NewUserHandler(users *service.UserService) *UserHandler {
	return &UserHandler{users: users}
}

func (handler *UserHandler) List(c *gin.Context) {
	var query model.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	response, err := handler.users.List(c.Request.Context(), query)
	SendResponse(c, err, response)
}

func (handler *UserHandler) Create(c *gin.Context) {
	var form model.UserCreateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	user, err := handler.users.Create(c.Request.Context(), form)
	SendResponse(c, err, user)
}

func (handler *UserHandler) Update(c *gin.Context) {
	var form model.UserUpdateForm
	if err := c.ShouldBindJSON(&form); err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	err := handler.users.Update(c.Request.Context(), form)
	SendResponse(c, err, nil)
}
