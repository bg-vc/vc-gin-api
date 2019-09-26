package handler

import (
	"github.com/gin-gonic/gin"
	"vc-gin-api/dao"
	"vc-gin-api/model"
	"vc-gin-api/pkg/errno"
	"vc-gin-api/pkg/log"
)

func LoginAccount(c *gin.Context) {
	req := &model.AccountReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Errorf(err, "c.ShouldBindJSON error")
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	dao.NewAccountDao()
}
