package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type GetAllCategories struct{}

type GetAllCategoriesHandler CommandHandler[GetAllCategories, []*domain.Category]

type getAllCategoriesHandler struct {
	facade CategoryFacade
}

func NewGetAllCategoriesHandler(facade CategoryFacade, logger *slog.Logger) GetAllCategoriesHandler {
	return ApplyHandlerDecorators(getAllCategoriesHandler{facade: facade}, logger)
}

func (h getAllCategoriesHandler) Handle(ctx context.Context, cmd GetAllCategories) ([]*domain.Category, error) {
	return h.facade.GetAllCategories(), nil
}
