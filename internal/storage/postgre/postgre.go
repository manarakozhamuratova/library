package postgre

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DialDatabase(ctx context.Context, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
