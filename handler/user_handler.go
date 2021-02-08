package handler

import (
	"github.com/gin-gonic/gin"
	"vc-gin-api/model"
	"vc-gin-api/pkg/errno"
	"vc-gin-api/pkg/log"
	"vc-gin-api/service"
)

func QueryUsers(c *gin.Context) {
	req := &model.UserReq{}
	if err := c.ShouldBind(req); err != nil {
		log.Errorf(err, "c.ShouldBind error")
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	if req.Start < 0 {
		req.Start = 0
	}
	if req.Limit > 50 {
		req.Limit = 50
	}

	service := service.NewUserService()

	resp, err := service.QueryUsers(req)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, resp)
}

func AddUser(c *gin.Context) {
	form := &model.UserForm{}
	if err := c.ShouldBindJSON(form); err != nil {
		SendResponse(c, err, nil)
		return
	}

	service := service.NewUserService()

	if err := service.AddUser(form); err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}

func UpdateUser(c *gin.Context) {
	form := &model.UserForm{}
	if err := c.ShouldBindJSON(form); err != nil {
		SendResponse(c, err, nil)
		return
	}

	service := service.NewUserService()

	if err := service.UpdateUser(form); err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}
