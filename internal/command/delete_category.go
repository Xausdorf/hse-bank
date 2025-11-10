package command

import (
	"context"
	slog "log/slog"
)

type DeleteCategory struct {
	ID string
}

type DeleteCategoryHandler CommandHandler[DeleteCategory, NoReturn]

type deleteCategoryHandler struct {
	facade CategoryFacade
}

func NewDeleteCategoryHandler(facade CategoryFacade, logger *slog.Logger) DeleteCategoryHandler {
	return ApplyHandlerDecorators(deleteCategoryHandler{facade: facade}, logger)
}

func (h deleteCategoryHandler) Handle(ctx context.Context, cmd DeleteCategory) (NoReturn, error) {
	return NoReturn{}, h.facade.DeleteCategory(cmd.ID)
}
