package service

import (
	"vc-gin-api/dao"
	"vc-gin-api/model"
	"vc-gin-api/pkg"
	"vc-gin-api/pkg/errno"
	"vc-gin-api/pkg/log"
	"vc-gin-api/util"
)

type AccountService struct {
	accountDao *dao.AccountDao
}

func NewAccountService(accountDao *dao.AccountDao) *AccountService {
	return &AccountService{accountDao}
}

func (service *AccountService) LoginAccount(req *model.AccountReq) (*model.AccountResp, error) {
	account, err := service.accountDao.GetAccountByName(req.Name)
	if err != nil {
		return nil, errno.InternalServerError
	}
	if account.Name == "" {
		return nil, errno.ErrAccount
	}
	salt := account.Salt
	pwdCrypt := util.MD5(salt + req.Password)
	if pwdCrypt != account.PwdCrypt {
		return nil, errno.ErrAccount
	}

	accountType := 2
	if account.IsAdmin == 1 {
		accountType = 1
	}
	token, err := signApiToken(account.ID, account.Name, accountType)
	if err != nil {
		return nil, errno.InternalServerError
	}
	resp := &model.AccountResp{
		Name:  account.Name,
		Token: token,
	}

	return resp, nil
}

func signApiToken(id uint64, username string, accountType int) (string, error) {
	token, err := pkg.Sign(pkg.Context{ID: id, Username: username}, accountType)
	if err != nil {
		log.Errorf(err, "signApiToken error")
		return "", nil
	}

	return token, nil
}
