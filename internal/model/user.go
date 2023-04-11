package model

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" swaggerignore:"true"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Wallet    int       `json:"wallet"`
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
	ID            uint   `json:"-" swaggerignore:"true"`
	OldPassword   string `json:"old_password"`
	NewPassword   string `json:"new_password"`
	ReNewPassword string `json:"re_new_password"`
}

type UserListing struct {
	ID            uint           `gorm:"primarykey" swaggerignore:"true"`
	Username      string         `json:"username" gorm:"unique"`
	BorrowedBooks []BorrowedBook `json:"borrowed_books" gorm:"-"`
}

type UserListingBookCount struct {
	ID       uint   `gorm:"primarykey" swaggerignore:"true"`
	Username string `json:"username" gorm:"unique"`
	Count    uint   `json:"count" gorm:"-"`
}

type contextKey string

var (
	ContextUsername = contextKey("username")
)
