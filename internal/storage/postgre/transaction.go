package postgre

import (
	"context"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: DB}
}

func (r *TransactionRepository) Create(ctx context.Context, tr *model.Transaction) error {
	result := r.DB.WithContext(ctx).Create(&tr)
	return result.Error
}
