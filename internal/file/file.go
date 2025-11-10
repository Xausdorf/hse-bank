package file

import "github.com/Xausdorf/hse-bank/internal/domain"

type FileData struct {
	Accounts   []domain.BankAccountDTO
	Categories []domain.CategoryDTO
	Operations []domain.OperationDTO
}
