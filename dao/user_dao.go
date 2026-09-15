package dao

import (
	"context"
	"database/sql"

	"github.com/bg-vc/vc-gin-api/model"
	"github.com/jmoiron/sqlx"
)

type UserDAO struct {
	db *sqlx.DB
}

func NewUserDAO(db *sqlx.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) List(ctx context.Context, query model.UserQuery) ([]model.User, int64, error) {
	const filter = ` WHERE (? = '' OR name = ?) AND (? = '' OR address LIKE CONCAT('%', ?, '%')) AND (? = '' OR email = ?)`
	args := []any{query.Name, query.Name, query.Address, query.Address, query.Email, query.Email}

	var total int64
	if err := dao.db.GetContext(ctx, &total, `SELECT COUNT(*) FROM user`+filter, args...); err != nil {
		return nil, 0, err
	}

	users := make([]model.User, 0)
	listArgs := append(args, query.Limit, query.Start)
	const columns = `SELECT id, name, age, address, email FROM user`
	if err := dao.db.SelectContext(ctx, &users, columns+filter+` ORDER BY id ASC LIMIT ? OFFSET ?`, listArgs...); err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (dao *UserDAO) Create(ctx context.Context, form model.UserCreateForm) (int64, error) {
	const query = `INSERT INTO user(name, age, address, email) VALUES (?, ?, ?, ?)`
	result, err := dao.db.ExecContext(ctx, query, form.Name, form.Age, form.Address, form.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (dao *UserDAO) Update(ctx context.Context, form model.UserUpdateForm) error {
	const query = `UPDATE user SET name = ?, age = ?, address = ?, email = ? WHERE id = ?`
	result, err := dao.db.ExecContext(ctx, query, form.Name, form.Age, form.Address, form.Email, form.ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
