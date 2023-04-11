package model

import "time"

type Transaction struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Sum       int  `gorm:"not null" json:"sum"`
	UserID    uint `gorm:"not null" json:"userID"`
	BookID    uint `gorm:"not null" json:"bookID"`
	CreatedAt time.Time
}
