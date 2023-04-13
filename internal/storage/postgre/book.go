package postgre

import (
	"context"
	"errors"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) Get(ctx context.Context, id uint) (model.Book, error) {
	var book model.Book
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *BookRepository) Create(ctx context.Context, book model.Book) (uint, error) {
	result := r.DB.WithContext(ctx).Create(&book)
	return book.ID, result.Error
}

func (r *BookRepository) BookIsAvailable(ctx context.Context, bookID uint) error {
	var bookBorrow model.BookBorrow
	err := r.DB.WithContext(ctx).Where("book_id = ?", bookID).Where("return_date is null").First(&bookBorrow).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if bookBorrow.ID != 0 {
		return errors.New("book already taken")
	}
	return nil
}

func (r *BookRepository) GetBookStatus(ctx context.Context, bookID, userID uint) (model.BookBorrow, error) {
	var bookBorrow model.BookBorrow
	err := r.DB.WithContext(ctx).Where("book_id = ?", bookID).Where("user_id = ?", userID).Where("return_date is null").First(&bookBorrow).Error
	return bookBorrow, err
}

func (r *BookRepository) Update(ctx context.Context, req model.BookBorrow) error {
	return r.DB.Where("id = ?", req.ID).Updates(&req).Error
}

func (r *BookRepository) TakeABook(ctx context.Context, req model.BookBorrow) error {
	result := r.DB.WithContext(ctx).Create(&req)
	return result.Error
}

func NewBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{DB: DB}
}
