package command

import (
	"context"
	slog "log/slog"
	"time"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type CreateOperation struct {
	AccountID   string
	CategoryID  string
	Amount      int64
	Date        time.Time
	Description string
}

type CreateOperationHandler CommandHandler[CreateOperation, *domain.Operation]

type createOperationHandler struct {
	facade OperationFacade
}

func NewCreateOperationHandler(facade OperationFacade, logger *slog.Logger) CreateOperationHandler {
	return ApplyHandlerDecorators(createOperationHandler{facade: facade}, logger)
}

func (h createOperationHandler) Handle(ctx context.Context, cmd CreateOperation) (*domain.Operation, error) {
	return h.facade.CreateOperation(cmd.AccountID, cmd.CategoryID, cmd.Amount, cmd.Date, cmd.Description)
}
