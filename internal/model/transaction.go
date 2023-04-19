package model

import "time"

type Transaction struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	Sum       float64 `gorm:"not null" json:"sum"`
	UserID    uint    `gorm:"not null" json:"userID"`
	BookID    uint    `gorm:"not null" json:"bookID"`
	Duration  uint    `gorm:"not null"`
	CreatedAt time.Time
}

type RentDuration struct {
	Duration uint `json:"duration"`
}
