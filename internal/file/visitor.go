package file

import "github.com/Xausdorf/hse-bank/internal/domain"

type marshalVisitor struct {
	fileData FileData
}

func (v *marshalVisitor) VisitBankAccount(account *domain.BankAccount) {
	if account == nil {
		return
	}
	accountDTO := domain.BankAccountDTO{
		ID:      account.ID(),
		Name:    account.Name(),
		Balance: account.Balance(),
	}
	v.fileData.Accounts = append(v.fileData.Accounts, accountDTO)
}

func (v *marshalVisitor) VisitCategory(category *domain.Category) {
	if category == nil {
		return
	}
	categoryDTO := domain.CategoryDTO{
		ID:            category.ID(),
		Name:          category.Name(),
		OperationType: category.OperationType().String(),
	}
	v.fileData.Categories = append(v.fileData.Categories, categoryDTO)
}

func (v *marshalVisitor) VisitOperation(operation *domain.Operation) {
	if operation == nil {
		return
	}
	operationDTO := domain.OperationDTO{
		ID:          operation.ID(),
		AccountID:   operation.AccountID(),
		CategoryID:  operation.CategoryID(),
		Amount:      operation.Amount(),
		Date:        operation.Date(),
		Description: operation.Description(),
	}
	v.fileData.Operations = append(v.fileData.Operations, operationDTO)
}
