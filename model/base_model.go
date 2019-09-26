package model

type Page struct {
	Start int64 `form:"start,default=0"`
	Limit int64 `form:"limit,default=20"`
}
