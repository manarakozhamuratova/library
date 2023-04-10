package model

import (
	"time"
)

type Book struct {
	ID        uint      `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	Name      string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
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
