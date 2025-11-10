package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
	"github.com/Xausdorf/hse-bank/internal/file"
)

type ExportAll struct {
	FilePath string
	Exporter *file.Exporter
}

type ExportAllHandler CommandHandler[ExportAll, NoReturn]

type exportAllHandler struct {
	acctFacade      BankAccountFacade
	categoryFacade  CategoryFacade
	operationFacade OperationFacade
}

func NewExportAllHandler(acct BankAccountFacade, cat CategoryFacade, op OperationFacade, logger *slog.Logger) ExportAllHandler {
	return ApplyHandlerDecorators(exportAllHandler{acctFacade: acct, categoryFacade: cat, operationFacade: op}, logger)
}

func (h exportAllHandler) Handle(ctx context.Context, cmd ExportAll) (NoReturn, error) {
	accounts := h.acctFacade.GetAllBankAccounts()
	categories := h.categoryFacade.GetAllCategories()
	operations, err := h.operationFacade.GetOperationsWithFilter(func(_ *domain.Operation) bool { return true })
	if err != nil {
		return NoReturn{}, err
	}

	if err := cmd.Exporter.Export(cmd.FilePath, accounts, categories, operations); err != nil {
		return NoReturn{}, err
	}
	return NoReturn{}, nil
}
