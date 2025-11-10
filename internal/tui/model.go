package tui

import (
	slog "log/slog"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Xausdorf/hse-bank/internal/service"
)

type screen int

const (
	screenMain screen = iota
	screenEntityMenu
	screenForm
	screenResult
)

type Model struct {
	svc    *service.Service
	logger *slog.Logger

	screen screen

	mainChoices []string
	cursor      int

	entityChoices []string
	entity        string

	prompts  []string
	answers  []string
	curField int
	onSubmit func([]string) (string, error)
	ti       textinput.Model

	result string
}

func NewModel(svc *service.Service, logger *slog.Logger) *Model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.CharLimit = 256
	ti.Width = 60
	ti.Focus()

	return &Model{
		svc:           svc,
		logger:        logger,
		screen:        screenMain,
		mainChoices:   []string{"BankAccount", "Category", "Operation", "File", "Quit"},
		entityChoices: []string{"View by ID", "Create", "Delete by ID", "Update", "View all", "Back"},
		ti:            ti,
	}
}

func (m *Model) Init() tea.Cmd { return nil }
