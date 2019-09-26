package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vc-gin-api/model"
	"vc-gin-api/pkg/log"
)

type AccountDao struct {
	dbRead  *sqlx.DB
	dbWrite *sqlx.DB
}

func NewAccountDao(dbRead, dbWrite *sqlx.DB) *AccountDao {
	return &AccountDao{dbRead, dbWrite}
}

func (dao *AccountDao) GetAccountByName(name string) (*model.Account, error) {
	sql := fmt.Sprintf(`select id, name, salt, pwd_crypt, status, is_admin `)
	sql += fmt.Sprintf(`form account `)
	sql += fmt.Sprintf(`where name=? `)
	log.Infof("GetAccountByName sql:%v", sql)

	rows, err := dao.dbRead.Queryx(sql, name)
	if err != nil {
		log.Errorf(err, "GetAccountByName error")
		return nil, err
	}
	defer rows.Close()
	item := &model.Account{}
	for rows.Next() {
		rows.StructScan(item)
	}

	return item, nil
}

func (dao *AccountDao) UpdateAccountPwd(name, salt, pwdCrypt string) error {
	sql := fmt.Sprintf(`update account set salt=?, pwd_crypt=? where name=? `)
	if _, err := dao.dbWrite.Exec(sql, salt, pwdCrypt, name); err != nil {
		log.Errorf(err, "UpdateAccountPwd error")
		return err
	}
	return nil
}
