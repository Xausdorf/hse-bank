package file

import (
	"fmt"
	"os"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type MarshalVisitor interface {
	VisitBankAccount(account *domain.BankAccount)
	VisitCategory(category *domain.Category)
	VisitOperation(operation *domain.Operation)
	BuildData() ([]byte, error)
}

type Exporter struct {
	visitor MarshalVisitor
}

func NewExporter(visitor MarshalVisitor) *Exporter {
	return &Exporter{visitor: visitor}
}

func (e *Exporter) Export(
	filePath string,
	accounts []*domain.BankAccount,
	categories []*domain.Category,
	operations []*domain.Operation,
) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create/open file: %w", err)
	}
	defer file.Close()

	for _, account := range accounts {
		e.visitor.VisitBankAccount(account)
	}

	for _, category := range categories {
		e.visitor.VisitCategory(category)
	}

	for _, operation := range operations {
		e.visitor.VisitOperation(operation)
	}

	bytes, err := e.visitor.BuildData()
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	if _, err = file.Write(bytes); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
