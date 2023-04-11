package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"github.com/manarakozhamuratova/one-lab-task2/internal/storage"
)

type IBookService interface {
	Create(ctx context.Context, book model.Book) (uint, error)
	TakeABook(ctx context.Context, op model.BookOperation) error
	GiveTheBook(ctx context.Context, op model.BookOperation) error
	BuyABook(ctx context.Context, tr model.Transaction) error
}

var _ IBookService = (*BookService)(nil)

type BookService struct {
	repo *storage.Storage
}

func NewBookService(repo *storage.Storage) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) Create(ctx context.Context, book model.Book) (uint, error) {
	return b.repo.Book.Create(ctx, book)
}

func (b *BookService) TakeABook(ctx context.Context, op model.BookOperation) error {
	if _, err := b.repo.Book.Get(ctx, op.BookID); err != nil {
		return err
	}
	if err := b.repo.Book.BookIsAvailable(ctx, op.BookID); err != nil {
		return fmt.Errorf("cannot take the book: %w", err)
	}
	req := model.BookBorrow{
		UserID: op.UserID,
		BookID: op.BookID,
	}
	if err := b.repo.Book.TakeABook(ctx, req); err != nil {
		return err
	}
	return nil
}

func (b *BookService) GiveTheBook(ctx context.Context, op model.BookOperation) error {
	if _, err := b.repo.Book.Get(ctx, op.BookID); err != nil {
		return err
	}
	bookStatus, err := b.repo.Book.GetBookStatus(ctx, op.BookID, op.UserID)
	if err != nil {
		return err
	}
	t := time.Now().UTC()
	bookStatus.ReturnDate = &t
	return b.repo.Book.Update(ctx, bookStatus)
}

func (b *BookService) BuyABook(ctx context.Context, tr model.Transaction) error {
	user, err := b.repo.User.Get(ctx, tr.UserID)
	if err != nil {
		return err
	}
	book, err := b.repo.Book.Get(ctx, tr.BookID)
	if err != nil {
		return err
	}
	if user.Wallet < book.Price {
		return errors.New("not enough money")
	}
	tr.Sum = book.Price
	user.Wallet = user.Wallet - book.Price
	if err := b.repo.User.Update(ctx, user); err != nil {
		return err
	}
	return b.repo.Transaction.Create(ctx, &tr)
}
