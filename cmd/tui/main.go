package main

import (
	"fmt"
	slog "log/slog"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/dig"

	cmd "github.com/Xausdorf/hse-bank/internal/command"
	"github.com/Xausdorf/hse-bank/internal/facade"
	"github.com/Xausdorf/hse-bank/internal/factory"
	inm "github.com/Xausdorf/hse-bank/internal/repository/inmemory"
	"github.com/Xausdorf/hse-bank/internal/service"
	"github.com/Xausdorf/hse-bank/internal/tui"
)

func setupContainer() (*dig.Container, error) {
	c := dig.New()

	c.Provide(func() (*slog.Logger, error) {
		f, err := os.OpenFile("hse-bank-tui.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		handler := slog.NewTextHandler(f, &slog.HandlerOptions{AddSource: false})
		logger := slog.New(handler)
		return logger, nil
	})

	if err := c.Provide(factory.NewBankAccountFactory, dig.As(new(facade.BankAccountFactory))); err != nil {
		return nil, fmt.Errorf("failed to provide account factory: %w", err)
	}
	if err := c.Provide(factory.NewCategoryFactory, dig.As(new(facade.CategoryFactory))); err != nil {
		return nil, fmt.Errorf("failed to provide category factory: %w", err)
	}
	if err := c.Provide(factory.NewOperationFactory, dig.As(new(facade.OperationFactory))); err != nil {
		return nil, fmt.Errorf("failed to provide operation factory: %w", err)
	}

	if err := c.Provide(inm.NewBankAccountRepository, dig.As(new(facade.BankAccountRepository))); err != nil {
		return nil, fmt.Errorf("failed to provide account repository: %w", err)
	}
	if err := c.Provide(inm.NewCategoryRepository, dig.As(new(facade.CategoryRepository))); err != nil {
		return nil, fmt.Errorf("failed to provide category repository: %w", err)
	}
	if err := c.Provide(inm.NewOperationRepository, dig.As(new(facade.OperationRepository))); err != nil {
		return nil, fmt.Errorf("failed to provide operation repository: %w", err)
	}

	if err := c.Provide(facade.NewBankAccountFacade, dig.As(new(cmd.BankAccountFacade))); err != nil {
		return nil, fmt.Errorf("failed to provide account facade: %w", err)
	}
	if err := c.Provide(facade.NewCategoryFacade, dig.As(new(cmd.CategoryFacade))); err != nil {
		return nil, fmt.Errorf("failed to provide category facade: %w", err)
	}
	if err := c.Provide(facade.NewOperationFacade, dig.As(new(cmd.OperationFacade))); err != nil {
		return nil, fmt.Errorf("failed to provide operation facade: %w", err)
	}

	return c, nil
}

func main() {
	c, err := setupContainer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup container: %v\n", err)
		os.Exit(1)
	}

	err = c.Invoke(func(accountFacade cmd.BankAccountFacade, categoryFacade cmd.CategoryFacade, operationFacade cmd.OperationFacade, acctRepo facade.BankAccountRepository, catRepo facade.CategoryRepository, opRepo facade.OperationRepository, logger *slog.Logger) {
		svc := service.NewService(accountFacade, categoryFacade, operationFacade, logger)

		p := tea.NewProgram(tui.NewModel(svc, logger), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			logger.Error("could not start TUI", "error", err)
			os.Exit(1)
		}

		time.Sleep(100 * time.Millisecond)
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start application: %v\n", err)
		os.Exit(1)
	}
}
