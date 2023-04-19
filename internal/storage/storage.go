package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/manarakozhamuratova/one-lab-task2/config"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"github.com/manarakozhamuratova/one-lab-task2/internal/storage/postgre"

	"gorm.io/gorm"
)

type IBookRepository interface {
	Create(ctx context.Context, book model.Book) (uint, error)
	Update(ctx context.Context, req model.BookBorrow) error
	TakeABook(ctx context.Context, req model.BookBorrow) error
	BookIsAvailable(ctx context.Context, bookID uint) error
	Get(ctx context.Context, id uint) (model.Book, error)
	GetBookStatus(ctx context.Context, bookID, userID uint) (model.BookBorrow, error)
}

var _ IBookRepository = (*postgre.BookRepository)(nil)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) (model.CreateResp, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, ID int) error
	GetByUsername(ctx context.Context, username string) (model.User, error)
	Get(ctx context.Context, id uint) (model.User, error)
	GetUsersWithActiveBorrowedBooks(ctx context.Context) ([]model.UserListing, error)
	GetUsersWithBorrowedBookCountByDate(ctx context.Context) ([]model.UserListingBookCount, error)
}

var _ IUserRepository = (*postgre.UserRepo)(nil)

type ITransactionRepository interface {
	CreateBuyTransaction(ctx context.Context, tr *model.Transaction) error
	ListRentedBooksRevenue(ctx context.Context) ([]model.RentedBook, error)
}

var _ ITransactionRepository = (*postgre.TransactionRepository)(nil)

type Storage struct {
	pg          *gorm.DB
	Book        IBookRepository
	User        IUserRepository
	Transaction ITransactionRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBPort)
	pgDB, err := postgre.DialDatabase(ctx, dsn)
	if err != nil {
		return nil, err
	}

	d, err := pgDB.DB()
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(d, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		cfg.DBMigrationsPath,
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	uRepo := postgre.NewUserRepo(pgDB)
	bRepo := postgre.NewBookRepository(pgDB)
	trRepo := postgre.NewTransactionRepository(pgDB)

	var storage Storage
	storage.User = uRepo
	storage.Book = bRepo
	storage.Transaction = trRepo
	return &storage, nil
}
