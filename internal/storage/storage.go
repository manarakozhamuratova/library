package storage

import "gorm.io/gorm"

type Storage struct {
	pg *gorm.DB
}
