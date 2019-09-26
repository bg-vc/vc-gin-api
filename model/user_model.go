package model

type User struct {
	ID      int64  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Age     int    `db:"age" json:"age"`
	Address string `db:"address" json:"address"`
}

type UserReq struct {
	Page
	Name    string `form:"name"`
	Address string `form:"address"`
}

type UserResp struct {
	Total int64   `json:"total"`
	Rows  []*User `json:"rows"`
}

type UserForm struct {
	Name    string `form:"name" binding:"required"`
	Age     int64  `form:"age" binding:"required"`
	Address string `form:"address" binding:"required"`
}
