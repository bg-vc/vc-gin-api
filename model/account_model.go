package model

type Account struct {
	ID       uint64 `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Salt     string `db:"salt" json:"salt"`
	PwdCrypt string `db:"pwd_crypt" json:"pwdCrypt"`
	Status   int    `db:"status" json:"status"`
	IsAdmin  int    `db:"is_admin" json:"isAdmin"`
}

type LoginForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AccountResp struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UpdatePwdForm struct {
	Name   string `form:"name" binding:"required"`
	OldPwd string `form:"oldPwd" binding:"required"`
	NewPwd string `form:"newPwd" binding:"required"`
}
