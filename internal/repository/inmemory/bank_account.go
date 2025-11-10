package inmemory

import (
	"errors"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type BankAccountRepository struct {
	accountByID map[string]*domain.BankAccount
}

func NewBankAccountRepository() *BankAccountRepository {
	return &BankAccountRepository{
		accountByID: make(map[string]*domain.BankAccount),
	}
}

func (r *BankAccountRepository) Save(account *domain.BankAccount) error {
	if account == nil {
		return errors.New("bank account is nil")
	}
	if _, ok := r.accountByID[account.ID()]; !ok {
		r.accountByID[account.ID()] = account
	}
	return nil
}

func (r *BankAccountRepository) GetByID(id string) (*domain.BankAccount, error) {
	account, ok := r.accountByID[id]
	if !ok {
		return nil, errors.New("bank account not found")
	}
	return account, nil
}

func (r *BankAccountRepository) Update(account *domain.BankAccount) error {
	if account == nil {
		return errors.New("bank account is nil")
	}
	if _, ok := r.accountByID[account.ID()]; !ok {
		return errors.New("bank account not found")
	}
	r.accountByID[account.ID()] = account
	return nil
}

func (r *BankAccountRepository) Delete(id string) error {
	if _, ok := r.accountByID[id]; !ok {
		return errors.New("bank account not found")
	}
	delete(r.accountByID, id)
	return nil
}

func (r *BankAccountRepository) GetAll() []*domain.BankAccount {
	accounts := make([]*domain.BankAccount, 0, len(r.accountByID))
	for _, account := range r.accountByID {
		accounts = append(accounts, account)
	}
	return accounts
}
