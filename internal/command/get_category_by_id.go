package command

import (
	"context"
	slog "log/slog"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type GetCategoryByID struct {
	ID string
}

type GetCategoryByIDHandler CommandHandler[GetCategoryByID, *domain.Category]

type getCategoryByIDHandler struct {
	facade CategoryFacade
}

func NewGetCategoryByIDHandler(facade CategoryFacade, logger *slog.Logger) GetCategoryByIDHandler {
	return ApplyHandlerDecorators(getCategoryByIDHandler{facade: facade}, logger)
}

func (h getCategoryByIDHandler) Handle(ctx context.Context, cmd GetCategoryByID) (*domain.Category, error) {
	return h.facade.GetCategoryByID(cmd.ID)
}
