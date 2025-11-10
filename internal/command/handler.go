package command

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type NoReturn struct{}

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context, command C) (R, error)
}

type loggingDecorator[C any, R any] struct {
	handler CommandHandler[C, R]
	logger  *slog.Logger
}

func (d *loggingDecorator[C, R]) Handle(ctx context.Context, command C) (R, error) {
	start := time.Now()
	d.logger.Info("handling command", "command", fmt.Sprintf("%T", command))

	result, err := d.handler.Handle(ctx, command)

	dur := time.Since(start)
	if err != nil {
		d.logger.Error("command failed", "command", fmt.Sprintf("%T", command), "duration", dur.String(), "error", err)
		return result, err
	}
	d.logger.Info("command handled", "command", fmt.Sprintf("%T", command), "duration", dur.String())
	return result, nil
}

func ApplyHandlerDecorators[C any, R any](handler CommandHandler[C, R], logger *slog.Logger) CommandHandler[C, R] {
	return &loggingDecorator[C, R]{handler: handler, logger: logger}
}
