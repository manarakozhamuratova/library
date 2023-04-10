package model

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" swaggerignore:"true"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `swaggerignore:"true"`
}

type CreateResp struct {
	ID        uint
	CreatedAt time.Time
}

type AuthUser struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type UpdatePassword struct {
	Username      string `json:"-" swaggerignore:"true"`
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	ReNewPassword string `json:"re_new_password"`
}

type contextKey string

var (
	ContextUsername = contextKey("username")
)
