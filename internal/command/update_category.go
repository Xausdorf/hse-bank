package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type UpdateCategory struct {
	ID     string
	Name   string
	OpType domain.OperationType
}

type UpdateCategoryHandler CommandHandler[UpdateCategory, NoReturn]

type updateCategoryHandler struct {
	facade CategoryFacade
}

func NewUpdateCategoryHandler(facade CategoryFacade, logger *slog.Logger) UpdateCategoryHandler {
	return ApplyHandlerDecorators(updateCategoryHandler{facade: facade}, logger)
}

func (h updateCategoryHandler) Handle(ctx context.Context, cmd UpdateCategory) (NoReturn, error) {
	return NoReturn{}, h.facade.UpdateCategory(cmd.ID, cmd.Name, cmd.OpType)
}
