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

func NewAccountService() *AccountService {
	return &AccountService{dao.NewAccountDao()}
}

func (service *AccountService) LoginAccount(form *model.LoginForm) (*model.AccountResp, *errno.Errno) {
	account, err := service.accountDao.GetAccountByName(form.Name)
	if err != nil {
		return nil, errno.InternalServerError
	}
	if account.Name == "" {
		return nil, errno.ErrAccount
	}
	salt := account.Salt
	pwdCrypt := util.MD5(salt + form.Password)
	if pwdCrypt != account.PwdCrypt {
		return nil, errno.ErrAccount
	}

	token, err := signApiToken(account.ID, account.Name)
	if err != nil {
		return nil, errno.InternalServerError
	}
	resp := &model.AccountResp{
		Name:  account.Name,
		Token: token,
	}

	return resp, nil
}

func (service *AccountService) UpdateAccountPwd(form *model.UpdatePwdForm) *errno.Errno {
	account, err := service.accountDao.GetAccountByName(form.Name)
	if err != nil {
		return errno.InternalServerError
	}
	if account.Name == "" {
		return errno.ErrAccount
	}
	oldSalt := account.Salt
	oldPwdCrypt := util.MD5(oldSalt + form.OldPwd)
	if account.PwdCrypt != oldPwdCrypt {
		return errno.ErrAccount
	}
	newSalt := util.GetRandomString(8)
	newPwdCrypt := util.MD5(newSalt + form.NewPwd)
	if err := service.accountDao.UpdateAccountPwd(form.Name, newSalt, newPwdCrypt); err != nil {
		return errno.InternalServerError
	}
	return nil
}

func signApiToken(id uint64, username string) (string, error) {
	token, err := pkg.Sign(pkg.Context{ID: id, Username: username})
	if err != nil {
		log.Errorf(err, "signApiToken error")
		return "", nil
	}

	return token, nil
}
