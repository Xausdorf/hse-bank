package command

import (
	"context"
	"errors"
	slog "log/slog"
	"strings"

	"github.com/Xausdorf/hse-bank/internal/domain"
	"github.com/Xausdorf/hse-bank/internal/file"
)

type Import struct {
	FilePath string
	Importer *file.ImporterImpl
}

type ImportHandler CommandHandler[Import, NoReturn]

type importHandler struct {
	acctFacade      BankAccountFacade
	categoryFacade  CategoryFacade
	operationFacade OperationFacade
}

func NewImportHandler(acctFacade BankAccountFacade, categoryFacade CategoryFacade, operationFacade OperationFacade, logger *slog.Logger) ImportHandler {
	return ApplyHandlerDecorators(importHandler{acctFacade: acctFacade, categoryFacade: categoryFacade, operationFacade: operationFacade}, logger)
}

func (h importHandler) Handle(ctx context.Context, cmd Import) (NoReturn, error) {
	accounts, categories, operations, err := cmd.Importer.Import(cmd.FilePath)
	if err != nil {
		return NoReturn{}, err
	}

	for _, a := range accounts {
		if _, err := h.acctFacade.CreateBankAccountWithID(a.ID, a.Name, a.Balance); err != nil {
			return NoReturn{}, err
		}
	}

	for _, c := range categories {
		var opType domain.OperationType
		switch strings.ToLower(c.OperationType) {
		case "income":
			opType = domain.Income
		case "expense":
			opType = domain.Expense
		default:
			return NoReturn{}, errors.New("unkown operation type")
		}
		if _, err := h.categoryFacade.CreateCategoryWithID(c.ID, c.Name, opType); err != nil {
			return NoReturn{}, err
		}
	}

	for _, o := range operations {
		if _, err := h.operationFacade.CreateOperationWithID(o.ID, o.AccountID, o.CategoryID, o.Amount, o.Date, o.Description); err != nil {
			return NoReturn{}, err
		}
	}

	return NoReturn{}, nil
}
