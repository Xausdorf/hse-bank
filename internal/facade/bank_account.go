package facade

import (
	"github.com/Xausdorf/hse-bank/internal/domain"
)

type BankAccountFacade struct {
	factory BankAccountFactory
	repo    BankAccountRepository
}

func (f *BankAccountFacade) CreateBankAccount(name string) (*domain.BankAccount, error) {
	bankAccount, err := f.factory.Create(name, 0)
	if err != nil {
		return nil, err
	}
	if err := f.repo.Save(bankAccount); err != nil {
		return nil, err
	}
	return bankAccount, nil
}

func (f *BankAccountFacade) UpdateBankAccount(id, name string, balance int64) error {
	bankAccount, err := f.factory.CreateWithID(id, name, balance)
	if err != nil {
		return err
	}
	if err := f.repo.Update(bankAccount); err != nil {
		return err
	}
	return nil
}

func (f *BankAccountFacade) GetBankAccountByID(id string) (*domain.BankAccount, error) {
	bankAccount, err := f.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return bankAccount, nil
}

func (f *BankAccountFacade) DeleteBankAccount(id string) error {
	if err := f.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (f *BankAccountFacade) GetAllBankAccounts() []*domain.BankAccount {
	return f.repo.GetAll()
}

func NewBankAccountFacade(factory BankAccountFactory, repo BankAccountRepository) *BankAccountFacade {
	return &BankAccountFacade{
		factory: factory,
		repo:    repo,
	}
}

func (f *BankAccountFacade) CreateBankAccountWithID(id, name string, balance int64) (*domain.BankAccount, error) {
	bankAccount, err := f.factory.CreateWithID(id, name, balance)
	if err != nil {
		return nil, err
	}
	if err := f.repo.Save(bankAccount); err != nil {
		return nil, err
	}
	return bankAccount, nil
}
