package command

import (
	"context"
	slog "log/slog"
)

type DeleteOperation struct {
	ID string
}

type DeleteOperationHandler CommandHandler[DeleteOperation, NoReturn]

type deleteOperationHandler struct {
	facade OperationFacade
}

func NewDeleteOperationHandler(facade OperationFacade, logger *slog.Logger) DeleteOperationHandler {
	return ApplyHandlerDecorators(deleteOperationHandler{facade: facade}, logger)
}

func (h deleteOperationHandler) Handle(ctx context.Context, cmd DeleteOperation) (NoReturn, error) {
	return NoReturn{}, h.facade.DeleteOperation(cmd.ID)
}
