package service

import (
	"errors"

	"github.com/manarakozhamuratova/one-lab-task2/internal/storage"
)

type Manager struct {
	// Book IBook
	User IUserService
}

var ErrEmptyStorage = errors.New("no storage provided")

func NewManager(storage *storage.Storage) (*Manager, error) {
	if storage == nil {
		return nil, ErrEmptyStorage
	}
	uSrv := NewUserService(storage)

	// bSrv := NewBookService(storage, *uSrv)

	return &Manager{
		// Book: bSrv,
		User: uSrv,
	}, nil
}
