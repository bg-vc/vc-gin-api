package handler

import (
	"github.com/gin-gonic/gin"
	"vc-gin-api/dao"
	"vc-gin-api/model"
	"vc-gin-api/pkg"
	"vc-gin-api/pkg/errno"
	"vc-gin-api/pkg/log"
	"vc-gin-api/service"
	"vc-gin-api/util"
)

func LoginAccount(c *gin.Context) {
	form := &model.LoginForm{}
	if err := c.ShouldBindJSON(form); err != nil {
		log.Errorf(err, "c.ShouldBindJSON error")
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	dao := dao.NewAccountDao(pkg.DBRead, pkg.DBWrite)
	service := service.NewAccountService(dao)

	resp, err := service.LoginAccount(form)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, resp)
}

func UpdateAccountPwd(c *gin.Context) {
	model := &model.UpdatePwdForm{}
	if err := c.ShouldBindJSON(model); err != nil {
		log.Errorf(err, "c.ShouldBindJSON error")
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	if util.IsHan(model.NewPwd) {
		SendResponse(c, errno.ErrPwdHasHan, nil)
		return
	}
	if len(model.NewPwd) < 8 {
		SendResponse(c, errno.ErrNewPwdLen, nil)
		return
	}

	dao := dao.NewAccountDao(pkg.DBRead, pkg.DBWrite)
	service := service.NewAccountService(dao)

	if err := service.UpdateAccountPwd(model); err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}
