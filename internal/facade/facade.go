package facade

import (
	"time"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type BankAccountFactory interface {
	Create(name string, balance int64) (*domain.BankAccount, error)
	CreateWithID(id, name string, balance int64) (*domain.BankAccount, error)
}

type CategoryFactory interface {
	Create(name string, operationType domain.OperationType) (*domain.Category, error)
	CreateWithID(id, name string, operationType domain.OperationType) (*domain.Category, error)
}

type OperationFactory interface {
	Create(accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error)
	CreateWithID(id, accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error)
}

type BankAccountRepository interface {
	Save(account *domain.BankAccount) error
	GetByID(id string) (*domain.BankAccount, error)
	Update(account *domain.BankAccount) error
	Delete(id string) error
	GetAll() []*domain.BankAccount
}

type CategoryRepository interface {
	Save(category *domain.Category) error
	GetByID(id string) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id string) error
	GetAll() []*domain.Category
}

type OperationRepository interface {
	Save(operation *domain.Operation) error
	GetByID(id string) (*domain.Operation, error)
	Update(operation *domain.Operation) error
	Delete(id string) error
	GetAll() []*domain.Operation
}
