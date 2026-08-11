package model

type User struct {
	ID      int64  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Age     int    `db:"age" json:"age"`
	Address string `db:"address" json:"address"`
}

type UserQuery struct {
	Page
	Name    string `form:"name"`
	Address string `form:"address"`
}

type UserResponse struct {
	Total int64  `json:"total"`
	Rows  []User `json:"rows"`
}

type UserCreateForm struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required,min=1,max=150"`
	Address string `json:"address" binding:"required"`
}

type UserUpdateForm struct {
	ID      int64  `json:"id" binding:"required,min=1"`
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required,min=1,max=150"`
	Address string `json:"address" binding:"required"`
}
