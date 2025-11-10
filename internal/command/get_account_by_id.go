package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type GetBankAccountByID struct {
	ID string
}

type GetBankAccountByIDHandler CommandHandler[GetBankAccountByID, *domain.BankAccount]

type getBankAccountByIDHandler struct {
	facade BankAccountFacade
}

func NewGetBankAccountByIDHandler(facade BankAccountFacade, logger *slog.Logger) GetBankAccountByIDHandler {
	return ApplyHandlerDecorators(getBankAccountByIDHandler{facade: facade}, logger)
}

func (h getBankAccountByIDHandler) Handle(ctx context.Context, cmd GetBankAccountByID) (*domain.BankAccount, error) {
	return h.facade.GetBankAccountByID(cmd.ID)
}
