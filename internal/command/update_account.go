package command

import (
	"context"
	slog "log/slog"
)

type UpdateBankAccount struct {
	ID      string
	Name    string
	Balance int64
}

type UpdateBankAccountHandler CommandHandler[UpdateBankAccount, NoReturn]

type updateBankAccountHandler struct {
	facade BankAccountFacade
}

func NewUpdateBankAccountHandler(facade BankAccountFacade, logger *slog.Logger) UpdateBankAccountHandler {
	return ApplyHandlerDecorators(updateBankAccountHandler{facade: facade}, logger)
}

func (h updateBankAccountHandler) Handle(ctx context.Context, cmd UpdateBankAccount) (NoReturn, error) {
	return NoReturn{}, h.facade.UpdateBankAccount(cmd.ID, cmd.Name, cmd.Balance)
}
