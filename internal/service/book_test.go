package service

import (
	"context"
	"testing"

	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
	"github.com/manarakozhamuratova/one-lab-task2/internal/storage"
	"github.com/manarakozhamuratova/one-lab-task2/internal/storage/mock"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
)

func TestBookService(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	st := &storage.Storage{
		Book:        mock.NewMockIBookRepository(ctrl),
		User:        mock.NewMockIUserRepository(ctrl),
		Transaction: mock.NewMockITransactionRepository(ctrl),
	}
	bookService := NewBookService(st)
	user := model.User{
		ID:       1,
		Username: "manara",
		Email:    "manara@gmail.com",
		Password: "123",
		Wallet:   200.0,
	}
	invalidUser := model.User{
		ID:       2,
		Username: "username",
		Email:    "user@mail.ru",
		Password: "852",
		Wallet:   50.0,
	}
	book := model.Book{
		ID:     1,
		Name:   "book",
		Author: "author",
		Price:  100.0,
	}

	testcases := []struct {
		name         string
		transactions model.Transaction
		mockFuncs    func()
		wantErr      bool
	}{
		{
			name:         "validTransactions",
			transactions: model.Transaction{UserID: user.ID, BookID: book.ID},
			mockFuncs: func() {
				st.User.(*mock.MockIUserRepository).EXPECT().Get(ctx, user.ID).Return(user, nil)
				st.Book.(*mock.MockIBookRepository).EXPECT().Get(ctx, book.ID).Return(book, nil)
				st.User.(*mock.MockIUserRepository).EXPECT().Update(ctx, gomock.Any()).Return(nil)
				st.Transaction.(*mock.MockITransactionRepository).EXPECT().Create(ctx, gomock.Any()).Return(nil)

			},

			wantErr: false,
		},
		{
			name:         "Not enough money",
			transactions: model.Transaction{UserID: invalidUser.ID, BookID: book.ID},
			mockFuncs: func() {
				st.User.(*mock.MockIUserRepository).EXPECT().Get(ctx, invalidUser.ID).Return(invalidUser, nil)
				st.Book.(*mock.MockIBookRepository).EXPECT().Get(ctx, book.ID).Return(book, nil)
			},
			wantErr: true,
		},
	}

	for _, test := range testcases {
		test.mockFuncs()
		err := bookService.BuyABook(ctx, test.transactions)
		if test.wantErr && err == nil {
			t.Errorf("expected error")
		} else if !test.wantErr && err != nil {
			t.Error("unexpected error")
		}
	}
}
