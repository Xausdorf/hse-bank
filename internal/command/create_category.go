package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type CreateCategory struct {
	Name   string
	OpType domain.OperationType
}

type CreateCategoryHandler CommandHandler[CreateCategory, *domain.Category]

type createCategoryHandler struct {
	facade CategoryFacade
}

func NewCreateCategoryHandler(facade CategoryFacade, logger *slog.Logger) CreateCategoryHandler {
	return ApplyHandlerDecorators(createCategoryHandler{facade: facade}, logger)
}

func (h createCategoryHandler) Handle(ctx context.Context, cmd CreateCategory) (*domain.Category, error) {
	return h.facade.CreateCategory(cmd.Name, cmd.OpType)
}
