package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type GetOperationByID struct {
	ID string
}

type GetOperationByIDHandler CommandHandler[GetOperationByID, *domain.Operation]

type getOperationByIDHandler struct {
	facade OperationFacade
}

func NewGetOperationByIDHandler(facade OperationFacade, logger *slog.Logger) GetOperationByIDHandler {
	return ApplyHandlerDecorators(getOperationByIDHandler{facade: facade}, logger)
}

func (h getOperationByIDHandler) Handle(ctx context.Context, cmd GetOperationByID) (*domain.Operation, error) {
	return h.facade.GetOperationByID(cmd.ID)
}
