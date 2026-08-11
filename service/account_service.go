package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bg-vc/vc-gin-api/dao"
	"github.com/bg-vc/vc-gin-api/model"
	"github.com/bg-vc/vc-gin-api/pkg"
	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	accounts *dao.AccountDAO
	tokens   *pkg.TokenManager
}

func NewAccountService(accounts *dao.AccountDAO, tokens *pkg.TokenManager) *AccountService {
	return &AccountService{accounts: accounts, tokens: tokens}
}

func (service *AccountService) Login(ctx context.Context, form model.LoginForm) (*model.AccountResponse, error) {
	account, err := service.accounts.GetByName(ctx, form.Name)
	if err != nil {
		return nil, err
	}
	if account == nil || account.Status != 1 {
		return nil, errno.ErrAccount
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(form.Password)); err != nil {
		return nil, errno.ErrAccount
	}

	token, expiresAt, err := service.tokens.Sign(account.ID, account.Name)
	if err != nil {
		return nil, err
	}
	return &model.AccountResponse{Name: account.Name, Token: token, ExpiresAt: expiresAt}, nil
}

func (service *AccountService) UpdatePassword(ctx context.Context, authenticatedName string, form model.UpdatePasswordForm) error {
	if form.Name != authenticatedName {
		return errno.ErrTokenInvalid
	}
	if len(form.NewPassword) < 8 {
		return errno.ErrPassword
	}

	account, err := service.accounts.GetByName(ctx, authenticatedName)
	if err != nil {
		return err
	}
	if account == nil {
		return errno.ErrAccount
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(form.OldPassword)); err != nil {
		return errno.ErrAccount
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(form.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := service.accounts.UpdatePassword(ctx, authenticatedName, string(hash)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errno.ErrDataNotExist
		}
		return err
	}
	return nil
}
