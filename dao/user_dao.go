package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vc-gin-api/model"
	"vc-gin-api/pkg"
	"vc-gin-api/pkg/log"
)

type UserDao struct {
	db *sqlx.DB
}

func NewUserDao() *UserDao {
	return &UserDao{pkg.DB}
}

func (dao *UserDao) GetUsers(req *model.UserReq) ([]*model.User, int64, error) {
	totalSql := fmt.Sprintf(`select count(1) as total from user `)
	querySql := fmt.Sprintf(`select id, name, age, address from user `)

	filterSql := fmt.Sprintf(`where 1=1 `)
	if req.Name != "" {
		filterSql += fmt.Sprintf(`and name='%v' `, req.Name)
	}
	if req.Address != "" {
		filterSql += fmt.Sprintf(`and address like '%v%v%v' `, "%", req.Address, "%")
	}

	sortSql := fmt.Sprintf(`order by id asc `)
	limitSql := fmt.Sprintf(`limit %v, %v `, req.Start, req.Limit)

	totalSql += filterSql
	querySql += filterSql + sortSql + limitSql

	log.Infof("GetUsers totalSql:%v", totalSql)
	row1, err := dao.db.Queryx(totalSql)
	if err != nil {
		log.Errorf(err, "GetUsers error:%v", totalSql)
		return nil, 0, err
	}
	defer row1.Close()
	total := int64(0)
	if row1.Next() {
		row1.Scan(&total)
	}

	log.Infof("GetUsers querySql:%v", querySql)
	row2, err := dao.db.Queryx(querySql)
	if err != nil {
		log.Errorf(err, "GetUsers error:%v", querySql)
		return nil, 0, err
	}
	defer row2.Close()
	list := make([]*model.User, 0)
	for row2.Next() {
		item := &model.User{}
		row2.StructScan(item)
		list = append(list, item)
	}

	return list, total, nil
}

func (dao *UserDao) AddUser(form *model.UserForm) error {
	sql := fmt.Sprintf(`insert into user(name, age, address) `)
	sql += fmt.Sprintf(`values(?, ?, ?) `)

	log.Infof("AddUser sql:%v", sql)
	if _, err := dao.db.Exec(sql, form.Name, form.Age, form.Address); err != nil {
		log.Errorf(err, "AddUser error:%v", sql)
		return err
	}
	return nil
}

func (dao *UserDao) UpdateUser(form *model.UserForm) error {
	sql := fmt.Sprintf(`update user set name=?, age=?, address=? `)
	sql += fmt.Sprintf(`where name=? `)

	log.Infof("UpdateUser sql:%v", sql)
	if _, err := dao.db.Exec(sql, form.Name, form.Age, form.Address, form.Name); err != nil {
		log.Errorf(err, "UpdateUser error:%v", sql)
	}
	return nil
}
