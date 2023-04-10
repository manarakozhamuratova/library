package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"github.com/manarakozhamuratova/one-lab-task2/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(ctx context.Context, req model.User) (model.CreateResp, error)
	CheckPassword(encPass, providedPassword string) error
	HashPassword(password string) (string, error)
	Auth(ctx context.Context, user model.AuthUser) (uint, error)
	UpdatePassword(ctx context.Context, user model.UpdatePassword) error
	GetUsersWithActiveBorrowedBooks(ctx context.Context) ([]model.UserListing, error)
	GetUsersWithBorrowedBookCountByDate(ctx context.Context) ([]model.UserListingBookCount, error)
}

var _ IUserService = (*UserService)(nil)

type UserService struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user model.User) (model.CreateResp, error) {
	var err error
	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		return model.CreateResp{}, err
	}
	return s.repo.User.Create(ctx, user)
}

func (s *UserService) CheckPassword(encPass, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encPass), []byte(providedPassword))
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *UserService) Auth(ctx context.Context, user model.AuthUser) (uint, error) {
	userFromDB, userErr := s.repo.User.GetByUsername(ctx, user.Username)
	if userErr != nil {
		return 0, userErr
	}
	checkErr := s.CheckPassword(userFromDB.Password, user.Password)
	if checkErr != nil {
		return 0, checkErr
	}
	return userFromDB.ID, nil
}

func (s *UserService) Update(ctx context.Context, user model.User) error {
	if user.ID == 0 && len(user.Username) == 0 {
		return fmt.Errorf("empty ID")
	}
	if len(user.Password) != 0 {
		user.Password, _ = s.HashPassword(user.Password)
	}
	return s.repo.User.Update(ctx, user)
}

func (s *UserService) UpdatePassword(ctx context.Context, upd model.UpdatePassword) error {
	if upd.NewPassword == upd.OldPassword {
		return errors.New("new password cannot be equal to old password")
	}
	if upd.NewPassword != upd.ReNewPassword {
		return errors.New("new password does not equal to new re password ")
	}
	user, err := s.repo.User.Get(ctx, upd.ID)
	if err != nil {
		return err
	}
	if err := s.CheckPassword(user.Password, upd.OldPassword); err != nil {
		return err
	}
	user.Password = upd.NewPassword
	return s.Update(ctx, user)
}

func (s *UserService) GetUsersWithActiveBorrowedBooks(ctx context.Context) ([]model.UserListing, error) {
	return s.repo.User.GetUsersWithActiveBorrowedBooks(ctx)
}

func (s *UserService) GetUsersWithBorrowedBookCountByDate(ctx context.Context) ([]model.UserListingBookCount, error) {
	return s.repo.User.GetUsersWithBorrowedBookCountByDate(ctx)
}
