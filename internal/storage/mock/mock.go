// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/manarakozhamuratova/one-lab-task2/internal/model"
)

// MockIBookRepository is a mock of IBookRepository interface.
type MockIBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBookRepositoryMockRecorder
}

// MockIBookRepositoryMockRecorder is the mock recorder for MockIBookRepository.
type MockIBookRepositoryMockRecorder struct {
	mock *MockIBookRepository
}

// NewMockIBookRepository creates a new mock instance.
func NewMockIBookRepository(ctrl *gomock.Controller) *MockIBookRepository {
	mock := &MockIBookRepository{ctrl: ctrl}
	mock.recorder = &MockIBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBookRepository) EXPECT() *MockIBookRepositoryMockRecorder {
	return m.recorder
}

// BookIsAvailable mocks base method.
func (m *MockIBookRepository) BookIsAvailable(ctx context.Context, bookID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookIsAvailable", ctx, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BookIsAvailable indicates an expected call of BookIsAvailable.
func (mr *MockIBookRepositoryMockRecorder) BookIsAvailable(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookIsAvailable", reflect.TypeOf((*MockIBookRepository)(nil).BookIsAvailable), ctx, bookID)
}

// Create mocks base method.
func (m *MockIBookRepository) Create(ctx context.Context, book model.Book) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIBookRepositoryMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBookRepository)(nil).Create), ctx, book)
}

// Get mocks base method.
func (m *MockIBookRepository) Get(ctx context.Context, id uint) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIBookRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIBookRepository)(nil).Get), ctx, id)
}

// GetBookStatus mocks base method.
func (m *MockIBookRepository) GetBookStatus(ctx context.Context, bookID, userID uint) (model.BookBorrow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookStatus", ctx, bookID, userID)
	ret0, _ := ret[0].(model.BookBorrow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookStatus indicates an expected call of GetBookStatus.
func (mr *MockIBookRepositoryMockRecorder) GetBookStatus(ctx, bookID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookStatus", reflect.TypeOf((*MockIBookRepository)(nil).GetBookStatus), ctx, bookID, userID)
}

// TakeABook mocks base method.
func (m *MockIBookRepository) TakeABook(ctx context.Context, req model.BookBorrow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeABook", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// TakeABook indicates an expected call of TakeABook.
func (mr *MockIBookRepositoryMockRecorder) TakeABook(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeABook", reflect.TypeOf((*MockIBookRepository)(nil).TakeABook), ctx, req)
}

// Update mocks base method.
func (m *MockIBookRepository) Update(ctx context.Context, req model.BookBorrow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIBookRepositoryMockRecorder) Update(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIBookRepository)(nil).Update), ctx, req)
}

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepository) Create(ctx context.Context, user model.User) (model.CreateResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(model.CreateResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepositoryMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), ctx, user)
}

// Delete mocks base method.
func (m *MockIUserRepository) Delete(ctx context.Context, ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserRepositoryMockRecorder) Delete(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserRepository)(nil).Delete), ctx, ID)
}

// Get mocks base method.
func (m *MockIUserRepository) Get(ctx context.Context, id uint) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUserRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUserRepository)(nil).Get), ctx, id)
}

// GetByUsername mocks base method.
func (m *MockIUserRepository) GetByUsername(ctx context.Context, username string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", ctx, username)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockIUserRepositoryMockRecorder) GetByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockIUserRepository)(nil).GetByUsername), ctx, username)
}

// GetUsersWithActiveBorrowedBooks mocks base method.
func (m *MockIUserRepository) GetUsersWithActiveBorrowedBooks(ctx context.Context) ([]model.UserListing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersWithActiveBorrowedBooks", ctx)
	ret0, _ := ret[0].([]model.UserListing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersWithActiveBorrowedBooks indicates an expected call of GetUsersWithActiveBorrowedBooks.
func (mr *MockIUserRepositoryMockRecorder) GetUsersWithActiveBorrowedBooks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersWithActiveBorrowedBooks", reflect.TypeOf((*MockIUserRepository)(nil).GetUsersWithActiveBorrowedBooks), ctx)
}

// GetUsersWithBorrowedBookCountByDate mocks base method.
func (m *MockIUserRepository) GetUsersWithBorrowedBookCountByDate(ctx context.Context) ([]model.UserListingBookCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersWithBorrowedBookCountByDate", ctx)
	ret0, _ := ret[0].([]model.UserListingBookCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersWithBorrowedBookCountByDate indicates an expected call of GetUsersWithBorrowedBookCountByDate.
func (mr *MockIUserRepositoryMockRecorder) GetUsersWithBorrowedBookCountByDate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersWithBorrowedBookCountByDate", reflect.TypeOf((*MockIUserRepository)(nil).GetUsersWithBorrowedBookCountByDate), ctx)
}

// Update mocks base method.
func (m *MockIUserRepository) Update(ctx context.Context, user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepositoryMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepository)(nil).Update), ctx, user)
}

// MockITransactionRepository is a mock of ITransactionRepository interface.
type MockITransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRepositoryMockRecorder
}

// MockITransactionRepositoryMockRecorder is the mock recorder for MockITransactionRepository.
type MockITransactionRepositoryMockRecorder struct {
	mock *MockITransactionRepository
}

// NewMockITransactionRepository creates a new mock instance.
func NewMockITransactionRepository(ctrl *gomock.Controller) *MockITransactionRepository {
	mock := &MockITransactionRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRepository) EXPECT() *MockITransactionRepositoryMockRecorder {
	return m.recorder
}

// CreateBuyTransaction mocks base method.
func (m *MockITransactionRepository) CreateBuyTransaction(ctx context.Context, tr *model.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuyTransaction", ctx, tr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBuyTransaction indicates an expected call of CreateBuyTransaction.
func (mr *MockITransactionRepositoryMockRecorder) CreateBuyTransaction(ctx, tr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuyTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).CreateBuyTransaction), ctx, tr)
}

// ListRentedBooksRevenue mocks base method.
func (m *MockITransactionRepository) ListRentedBooksRevenue(ctx context.Context) ([]model.RentedBook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRentedBooksRevenue", ctx)
	ret0, _ := ret[0].([]model.RentedBook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRentedBooksRevenue indicates an expected call of ListRentedBooksRevenue.
func (mr *MockITransactionRepositoryMockRecorder) ListRentedBooksRevenue(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRentedBooksRevenue", reflect.TypeOf((*MockITransactionRepository)(nil).ListRentedBooksRevenue), ctx)
}