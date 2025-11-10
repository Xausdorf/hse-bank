package domain

import "time"

type BankAccountDTO struct {
	ID      string
	Name    string
	Balance int64
}

type CategoryDTO struct {
	ID            string
	Name          string
	OperationType string
}

type OperationDTO struct {
	ID          string
	AccountID   string
	CategoryID  string
	Amount      int64
	Date        time.Time
	Description string
}
