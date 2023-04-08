package model

type User struct {
	ID       uint   `gorm:"primarykey" swaggerignore:"true"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
