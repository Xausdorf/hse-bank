package facade

import (
	"errors"
	"time"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

var (
	ErrInvalidAccountID    = errors.New("invalid account ID")
	ErrInvalidCategoryID   = errors.New("invalid category ID")
	ErrCannotUpdateBalance = errors.New("cannot update account balance")
)

type OperationFacade struct {
	factory        OperationFactory
	repo           OperationRepository
	accountFactory BankAccountFactory
	accountRepo    BankAccountRepository
	categoryRepo   CategoryRepository
}

func (f *OperationFacade) getAccount(accountID string) (*domain.BankAccount, error) {
	account, err := f.accountRepo.GetByID(accountID)
	if err != nil {
		return nil, ErrInvalidAccountID
	}
	return account, nil
}

func (f *OperationFacade) getCategory(categoryID string) (*domain.Category, error) {
	category, err := f.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, ErrInvalidCategoryID
	}
	return category, nil
}

func (f *OperationFacade) updateBalance(account *domain.BankAccount, amount int64, operationType domain.OperationType) error {
	if account == nil {
		return errors.New("account is nil")
	}
	if operationType == domain.Expense {
		amount = -amount
	}
	newAccount, err := f.accountFactory.CreateWithID(account.ID(), account.Name(), account.Balance())
	if err != nil {
		return err
	}
	newAccount.ApplyOperation(amount, operationType)
	if err := f.accountRepo.Update(newAccount); err != nil {
		return ErrCannotUpdateBalance
	}
	return nil
}

func (f *OperationFacade) cancelOperation(operationID string) error {
	operation, err := f.repo.GetByID(operationID)
	if err != nil {
		return err
	}
	account, err := f.getAccount(operation.AccountID())
	if err != nil {
		return err
	}
	category, err := f.getCategory(operation.CategoryID())
	if err != nil {
		return err
	}
	if err := f.repo.Delete(operationID); err != nil {
		return err
	}

	if err := f.updateBalance(account, operation.Amount(), domain.ReverseOperationType(category.OperationType())); err != nil {
		_ = f.repo.Save(operation)
		return err
	}
	return nil
}

func (f *OperationFacade) CreateOperation(accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error) {
	account, err := f.getAccount(accountID)
	if err != nil {
		return nil, err
	}
	category, err := f.getCategory(categoryID)
	if err != nil {
		return nil, err
	}

	operation, err := f.factory.Create(accountID, categoryID, amount, date, description)
	if err != nil {
		return nil, err
	}
	if err := f.repo.Save(operation); err != nil {
		return nil, err
	}

	if err := f.updateBalance(account, amount, category.OperationType()); err != nil {
		_ = f.repo.Delete(operation.ID())
		return nil, err
	}

	return operation, nil
}

func (f *OperationFacade) CreateOperationWithID(id, accountID, categoryID string, amount int64, date time.Time, description string) (*domain.Operation, error) {
	account, err := f.getAccount(accountID)
	if err != nil {
		return nil, err
	}
	category, err := f.getCategory(categoryID)
	if err != nil {
		return nil, err
	}

	operation, err := f.factory.CreateWithID(id, accountID, categoryID, amount, date, description)
	if err != nil {
		return nil, err
	}
	if err := f.repo.Save(operation); err != nil {
		return nil, err
	}

	if err := f.updateBalance(account, amount, category.OperationType()); err != nil {
		_ = f.repo.Delete(operation.ID())
		return nil, err
	}

	return operation, nil
}

func (f *OperationFacade) UpdateOperation(id, accountID, categoryID string, amount int64, date time.Time, description string) error {
	operation, err := f.factory.CreateWithID(id, accountID, categoryID, amount, date, description)
	if err != nil {
		return err
	}
	oldOperation, err := f.repo.GetByID(id)
	if err != nil {
		return err
	}
	oldAccount, err := f.getAccount(oldOperation.AccountID())
	if err != nil {
		return err
	}
	oldCategory, err := f.getCategory(oldOperation.CategoryID())
	if err != nil {
		return err
	}
	if err := f.cancelOperation(oldOperation.ID()); err != nil {
		return err
	}

	if err := func() error {
		account, err := f.getAccount(accountID)
		if err != nil {
			return err
		}
		category, err := f.getCategory(categoryID)
		if err != nil {
			return err
		}

		if err := f.repo.Save(operation); err != nil {
			return err
		}

		if err := f.updateBalance(account, amount, category.OperationType()); err != nil {
			_ = f.repo.Delete(id)
			return err
		}
		return nil
	}(); err != nil {
		_ = f.repo.Save(oldOperation)
		_ = f.updateBalance(oldAccount, oldOperation.Amount(), oldCategory.OperationType())
		return err
	}

	return nil
}

func (f *OperationFacade) GetOperationByID(id string) (*domain.Operation, error) {
	operation, err := f.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return operation, nil
}

func (f *OperationFacade) DeleteOperation(operationID string) error {
	if err := f.cancelOperation(operationID); err != nil {
		return err
	}
	return nil
}

func (f *OperationFacade) GetOperationsWithFilter(predicate func(*domain.Operation) bool) ([]*domain.Operation, error) {
	allOperations := f.repo.GetAll()
	var filteredOperations []*domain.Operation
	for _, op := range allOperations {
		if predicate(op) {
			filteredOperations = append(filteredOperations, op)
		}
	}
	return filteredOperations, nil
}

func NewOperationFacade(factory OperationFactory, repo OperationRepository, accountFactory BankAccountFactory, accountRepo BankAccountRepository, categoryRepo CategoryRepository) *OperationFacade {
	return &OperationFacade{
		factory:        factory,
		repo:           repo,
		accountFactory: accountFactory,
		accountRepo:    accountRepo,
		categoryRepo:   categoryRepo,
	}
}
