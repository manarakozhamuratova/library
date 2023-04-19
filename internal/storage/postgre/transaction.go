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

func (r *TransactionRepository) CreateBuyTransaction(ctx context.Context, tr *model.Transaction) error {
	result := r.DB.WithContext(ctx).Create(&tr)
	return result.Error
}

func (r *TransactionRepository) ListRentedBooksRevenue(ctx context.Context) ([]model.RentedBook, error) {
	var result []model.RentedBook
	err := r.DB.Table("transactions").
		Select("book_id as id, books.name as name, SUM(transactions.sum) as total").
		Where("duration != 0").
		Where(" NOW() < transactions.created_at + (duration || ' days')::INTERVAL").
		Group("book_id, books.name").
		Joins("JOIN books ON transactions.book_id = books.id").
		Scan(&result).
		Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
