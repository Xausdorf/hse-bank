package factory

import (
	"fmt"
	"time"

	"github.com/Xausdorf/hse-bank/internal/domain"

	"github.com/google/uuid"
)

type BankAccountFactory struct{}

func (f *BankAccountFactory) Create(name string, balance int64) (*domain.BankAccount, error) {
	id := uuid.New().String()
	return f.CreateWithID(id, name, balance)
}

func (f *BankAccountFactory) CreateWithID(id, name string, balance int64) (*domain.BankAccount, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if err := uuid.Validate(id); err != nil {
		return nil, fmt.Errorf("invalid account ID: %w", err)
	}
	return domain.NewBankAccount(id, name, balance), nil
}

type CategoryFactory struct{}

func (f *CategoryFactory) Create(name string, operationType domain.OperationType) (*domain.Category, error) {
	id := uuid.New().String()
	return f.CreateWithID(id, name, operationType)
}

func (f *CategoryFactory) CreateWithID(id, name string, operationType domain.OperationType) (*domain.Category, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if err := uuid.Validate(id); err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}
	return domain.NewCategory(id, name, operationType), nil
}

type OperationFactory struct{}

func (f *OperationFactory) Create(accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error) {
	id := uuid.New().String()
	return f.CreateWithID(id, accountID, categoryID, amount, date, description)
}

func (f *OperationFactory) CreateWithID(id, accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error) {
	if amount < 0 {
		return nil, ErrNegativeAmount
	}
	if err := uuid.Validate(id); err != nil {
		return nil, fmt.Errorf("invalid operation ID: %w", err)
	}
	if err := uuid.Validate(accountID); err != nil {
		return nil, fmt.Errorf("invalid account ID: %w", err)
	}
	if err := uuid.Validate(categoryID); err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}
	return domain.NewOperation(id, accountID, categoryID, amount, date, description), nil
}

func NewBankAccountFactory() *BankAccountFactory {
	return &BankAccountFactory{}
}

func NewCategoryFactory() *CategoryFactory {
	return &CategoryFactory{}
}

func NewOperationFactory() *OperationFactory {
	return &OperationFactory{}
}
