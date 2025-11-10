package command

import (
	"context"
	slog "log/slog"
)

type DeleteBankAccount struct {
	ID string
}

type DeleteBankAccountHandler CommandHandler[DeleteBankAccount, NoReturn]

type deleteBankAccountHandler struct {
	facade BankAccountFacade
}

func NewDeleteBankAccountHandler(facade BankAccountFacade, logger *slog.Logger) DeleteBankAccountHandler {
	return ApplyHandlerDecorators(deleteBankAccountHandler{facade: facade}, logger)
}

func (h deleteBankAccountHandler) Handle(ctx context.Context, cmd DeleteBankAccount) (NoReturn, error) {
	return NoReturn{}, h.facade.DeleteBankAccount(cmd.ID)
}
