package command

import (
	"context"
	slog "log/slog"
	"time"
)

type UpdateOperation struct {
	ID          string
	AccountID   string
	CategoryID  string
	Amount      int64
	Date        time.Time
	Description string
}

type UpdateOperationHandler CommandHandler[UpdateOperation, NoReturn]

type updateOperationHandler struct {
	facade OperationFacade
}

func NewUpdateOperationHandler(facade OperationFacade, logger *slog.Logger) UpdateOperationHandler {
	return ApplyHandlerDecorators(updateOperationHandler{facade: facade}, logger)
}

func (h updateOperationHandler) Handle(ctx context.Context, cmd UpdateOperation) (NoReturn, error) {
	return NoReturn{}, h.facade.UpdateOperation(cmd.ID, cmd.AccountID, cmd.CategoryID, cmd.Amount, cmd.Date, cmd.Description)
}
