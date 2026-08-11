package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bg-vc/vc-gin-api/dao"
	"github.com/bg-vc/vc-gin-api/model"
	"github.com/bg-vc/vc-gin-api/pkg/errno"
)

type UserService struct {
	users *dao.UserDAO
}

func NewUserService(users *dao.UserDAO) *UserService {
	return &UserService{users: users}
}

func (service *UserService) List(ctx context.Context, query model.UserQuery) (*model.UserResponse, error) {
	if query.Start < 0 {
		query.Start = 0
	}
	if query.Limit <= 0 {
		query.Limit = 20
	}
	if query.Limit > 50 {
		query.Limit = 50
	}
	users, total, err := service.users.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return &model.UserResponse{Rows: users, Total: total}, nil
}

func (service *UserService) Create(ctx context.Context, form model.UserCreateForm) (model.User, error) {
	id, err := service.users.Create(ctx, form)
	if err != nil {
		return model.User{}, err
	}
	return model.User{ID: id, Name: form.Name, Age: form.Age, Address: form.Address}, nil
}

func (service *UserService) Update(ctx context.Context, form model.UserUpdateForm) error {
	if err := service.users.Update(ctx, form); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errno.ErrDataNotExist
		}
		return err
	}
	return nil
}
