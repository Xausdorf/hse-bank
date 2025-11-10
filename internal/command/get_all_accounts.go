package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type GetAllBankAccounts struct{}

type GetAllBankAccountsHandler CommandHandler[GetAllBankAccounts, []*domain.BankAccount]

type getAllBankAccountsHandler struct {
	facade BankAccountFacade
}

func NewGetAllBankAccountsHandler(facade BankAccountFacade, logger *slog.Logger) GetAllBankAccountsHandler {
	return ApplyHandlerDecorators(getAllBankAccountsHandler{facade: facade}, logger)
}

func (h getAllBankAccountsHandler) Handle(ctx context.Context, cmd GetAllBankAccounts) ([]*domain.BankAccount, error) {
	return h.facade.GetAllBankAccounts(), nil
}
