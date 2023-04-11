package service

import (
	"errors"

	"github.com/manarakozhamuratova/one-lab-task2/internal/storage"
)

type Service struct {
	Book IBookService
	User IUserService
}

var ErrEmptyStorage = errors.New("no storage provided")

func NewService(storage *storage.Storage) (*Service, error) {
	if storage == nil {
		return nil, ErrEmptyStorage
	}
	uSrv := NewUserService(storage)
	bSrv := NewBookService(storage)
	return &Service{
		Book: bSrv,
		User: uSrv,
	}, nil
}
