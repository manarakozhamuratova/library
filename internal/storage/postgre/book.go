package postgre

import (
	"context"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) Create(ctx context.Context, book model.Book) (uint, error) {
	result := r.DB.WithContext(ctx).Omit("deleted_at").Create(&book)
	return book.ID, result.Error
}

func (r *BookRepository) Delete() {
	r.Get()
	//TODO implement me
	panic("implement me")
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}
