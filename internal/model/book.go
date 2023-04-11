package model

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Book struct {
	ID        uint      `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	Name      string    `gorm:"not null" validate:"required"`
	Author    string    `gorm:"not null" validate:"required"`
	Price     int       `gorm:"not null" validate:"required"`
	CreatedAt time.Time `swaggerignore:"true"`
}

type BookOperation struct {
	UserID    uint
	BookID    uint
	CreatedAt time.Time
}

type BookBorrow struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	BookID     uint      `gorm:"not null"`
	BorrowDate time.Time `gorm:"default:now()"`
	ReturnDate *time.Time
}

type BorrowedBook struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	BorrowDate time.Time `json:"borrow_date"`
}
