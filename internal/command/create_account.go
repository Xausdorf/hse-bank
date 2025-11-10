package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type CreateBankAccount struct {
	Name string
}

type CreateBankAccountHandler CommandHandler[CreateBankAccount, *domain.BankAccount]

type createBankAccountHandler struct {
	facade BankAccountFacade
}

func NewCreateBankAccountHandler(facade BankAccountFacade, logger *slog.Logger) CreateBankAccountHandler {
	return ApplyHandlerDecorators(createBankAccountHandler{facade: facade}, logger)
}

func (h createBankAccountHandler) Handle(ctx context.Context, cmd CreateBankAccount) (*domain.BankAccount, error) {
	return h.facade.CreateBankAccount(cmd.Name)
}
