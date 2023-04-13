package postgre

import (
	"context"
	"time"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (r *UserRepo) Create(ctx context.Context, user model.User) (model.CreateResp, error) {
	result := r.DB.WithContext(ctx).Omit("deleted_at").Create(&user)
	return model.CreateResp{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
	}, result.Error
}

func (r *UserRepo) Update(ctx context.Context, user model.User) error {
	var u model.User
	r.DB.Find(&u).Where("id = ?", user.ID)
	u.Wallet = user.Wallet
	return r.DB.Save(user).Error
}

func (r *UserRepo) Delete(ctx context.Context, ID int) error {
	return r.DB.WithContext(ctx).Delete(&model.User{}, ID).Error
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&res).Error
	return res, err
}

func (r *UserRepo) Get(ctx context.Context, id uint) (model.User, error) {
	var res model.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&res).Error
	return res, err
}

func (r *UserRepo) GetUsersWithActiveBorrowedBooks(ctx context.Context) ([]model.UserListing, error) {
	var users []model.UserListing

	// find all users
	if err := r.DB.WithContext(ctx).
		Table("users").
		Select("users.id as id, users.username as username").
		Scan(&users).Error; err != nil {
		return nil, err
	}

	// loop through users and get their active borrowed books
	for i := range users {
		var borrowedBooks []struct {
			BookID     uint
			BookName   string
			BorrowDate time.Time
		}

		if err := r.DB.WithContext(ctx).Table("book_borrows").
			Select("books.id as book_id, books.name as book_name, book_borrows.borrow_date as borrow_date").
			Joins("JOIN books ON book_borrows.book_id = books.id").
			Where("book_borrows.user_id = ? AND book_borrows.return_date IS NULL", users[i].ID).
			Scan(&borrowedBooks).Error; err != nil {
			return nil, err
		}

		for j := range borrowedBooks {
			users[i].BorrowedBooks = append(users[i].BorrowedBooks, model.BorrowedBook{
				ID:         borrowedBooks[j].BookID,
				Name:       borrowedBooks[j].BookName,
				BorrowDate: time.Now(),
			})
		}
	}

	return users, nil
}

func (r *UserRepo) GetUsersWithBorrowedBookCountByDate(ctx context.Context) ([]model.UserListingBookCount, error) {
	var users []model.UserListingBookCount

	// find all users
	if err := r.DB.WithContext(ctx).
		Table("users").
		Select("users.id as id, users.username as username").
		Scan(&users).Error; err != nil {
		return nil, err
	}

	// loop through users and get their borrowed book count by borrow date for the last 30 days
	for i := range users {
		var count uint

		thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

		if err := r.DB.WithContext(ctx).Table("book_borrows").
			Select("count(*)").
			Where("user_id = ? AND borrow_date >= ?", users[i].ID, thirtyDaysAgo).
			Group("date_trunc('day', borrow_date)").
			Order("date_trunc('day', borrow_date) DESC").
			Scan(&count).Error; err != nil {
			return nil, err
		}

		users[i].Count = count
	}

	return users, nil
}
