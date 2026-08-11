package dao

import (
	"context"
	"database/sql"

	"github.com/bg-vc/vc-gin-api/model"
	"github.com/jmoiron/sqlx"
)

type AccountDAO struct {
	db *sqlx.DB
}

func NewAccountDAO(db *sqlx.DB) *AccountDAO {
	return &AccountDAO{db: db}
}

func (dao *AccountDAO) GetByName(ctx context.Context, name string) (*model.Account, error) {
	const query = `SELECT id, name, password_hash, status, is_admin FROM account WHERE name = ?`
	account := &model.Account{}
	if err := dao.db.GetContext(ctx, account, query, name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return account, nil
}

func (dao *AccountDAO) UpdatePassword(ctx context.Context, name, passwordHash string) error {
	const query = `UPDATE account SET password_hash = ? WHERE name = ?`
	result, err := dao.db.ExecContext(ctx, query, passwordHash, name)
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
