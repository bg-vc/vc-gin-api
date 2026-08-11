package model

import "time"

type Account struct {
	ID           uint64 `db:"id"`
	Name         string `db:"name"`
	PasswordHash string `db:"password_hash"`
	Status       int    `db:"status"`
	IsAdmin      int    `db:"is_admin"`
}

type LoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountResponse struct {
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type UpdatePasswordForm struct {
	Name        string `json:"name" binding:"required"`
	OldPassword string `json:"oldPwd" binding:"required"`
	NewPassword string `json:"newPwd" binding:"required,min=8"`
}
