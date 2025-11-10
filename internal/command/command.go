package command

import (
	"time"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type BankAccountFacade interface {
	CreateBankAccount(name string) (*domain.BankAccount, error)
	CreateBankAccountWithID(id, name string, balance int64) (*domain.BankAccount, error)
	UpdateBankAccount(id, name string, balance int64) error
	GetBankAccountByID(id string) (*domain.BankAccount, error)
	DeleteBankAccount(id string) error
	GetAllBankAccounts() []*domain.BankAccount
}

type CategoryFacade interface {
	CreateCategory(name string, operationType domain.OperationType) (*domain.Category, error)
	CreateCategoryWithID(id, name string, operationType domain.OperationType) (*domain.Category, error)
	UpdateCategory(id, name string, opType domain.OperationType) error
	GetCategoryByID(id string) (*domain.Category, error)
	DeleteCategory(id string) error
	GetAllCategories() []*domain.Category
}

type OperationFacade interface {
	CreateOperation(accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error)
	CreateOperationWithID(id, accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error)
	UpdateOperation(id, accountID, categoryID string, amount int64, date time.Time, description string) error
	GetOperationByID(id string) (*domain.Operation, error)
	DeleteOperation(operationID string) error
	GetOperationsWithFilter(predicate func(*domain.Operation) bool) ([]*domain.Operation, error)
}
