package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type GetOperationsWithFilter struct {
	Predicate func(*domain.Operation) bool
}

type GetOperationsWithFilterHandler CommandHandler[GetOperationsWithFilter, []*domain.Operation]

type getOperationsWithFilterHandler struct {
	facade OperationFacade
}

func NewGetOperationsWithFilterHandler(facade OperationFacade, logger *slog.Logger) GetOperationsWithFilterHandler {
	return ApplyHandlerDecorators(getOperationsWithFilterHandler{facade: facade}, logger)
}

func (h getOperationsWithFilterHandler) Handle(ctx context.Context, cmd GetOperationsWithFilter) ([]*domain.Operation, error) {
	return h.facade.GetOperationsWithFilter(cmd.Predicate)
}
