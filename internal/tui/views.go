package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.screen {
	case screenMain:
		return m.updateMain(msg)
	case screenEntityMenu:
		return m.updateEntityMenu(msg)
	case screenForm:
		return m.updateForm(msg)
	case screenResult:
		return m.updateResult(msg)
	default:
		return m, nil
	}
}

func (m *Model) View() string {
	switch m.screen {
	case screenMain:
		return m.viewMain()
	case screenEntityMenu:
		return m.viewEntityMenu()
	case screenForm:
		return m.viewForm()
	case screenResult:
		return m.viewResult()
	default:
		return ""
	}
}

func (m *Model) updateMain(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.mainChoices)-1 {
				m.cursor++
			}
		case "enter":
			choice := m.mainChoices[m.cursor]
			if choice == "Quit" {
				return m, tea.Quit
			}
			m.entity = choice
			if choice == "File" {
				m.entityChoices = []string{"Export JSON", "Export YAML", "Import JSON", "Import YAML", "Back"}
			} else {
				m.entityChoices = []string{"View by ID", "Create", "Delete by ID", "Update", "View all", "Back"}
			}
			m.screen = screenEntityMenu
			m.cursor = 0
		}
	}
	return m, nil
}

func (m *Model) viewMain() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "Main menu - select entity (use ↑/↓, enter)\n\n")
	for i, choice := range m.mainChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		fmt.Fprintf(b, "%s %s\n", cursor, choice)
	}
	fmt.Fprintf(b, "\nPress q to quit")
	return b.String()
}

func (m *Model) updateEntityMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.screen = screenMain
			m.cursor = 0
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.entityChoices)-1 {
				m.cursor++
			}
		case "enter":
			choice := m.entityChoices[m.cursor]
			if choice == "Back" {
				m.screen = screenMain
				m.cursor = 0
				return m, nil
			}
			m.prepareFormFor(choice)
			m.screen = screenForm
			m.ti.Focus()
		}
	}
	return m, nil
}

func (m *Model) viewEntityMenu() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%s menu - choose action\n\n", m.entity)
	for i, choice := range m.entityChoices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		fmt.Fprintf(b, "%s %s\n", cursor, choice)
	}
	fmt.Fprintf(b, "\nPress esc to go back to main menu")
	return b.String()
}

func (m *Model) updateResult(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "esc":
			m.screen = screenEntityMenu
			m.cursor = 0
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) viewResult() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "Result:\n\n%s\n\n", m.result)
	fmt.Fprintf(b, "Press Enter or Esc to go back")
	return b.String()
}

func (m *Model) viewForm() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%s > %s\n\n", m.entity, m.prompts[m.curField])
	fmt.Fprintf(b, "%s\n\n", m.ti.View())
	fmt.Fprintf(b, "Enter - next / submit, Esc - back")
	return b.String()
}

func (m *Model) updateForm(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.ti, cmd = m.ti.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.screen = screenEntityMenu
			m.cursor = 0
			return m, nil
		case "enter":
			m.answers[m.curField] = m.ti.Value()
			if m.curField >= len(m.prompts)-1 {
				res, err := m.onSubmit(m.answers)
				if err != nil {
					m.result = "Error: " + err.Error()
				} else {
					m.result = res
				}
				m.screen = screenResult
				return m, nil
			}
			m.curField++
			m.ti.SetValue(m.answers[m.curField])
			m.ti.Placeholder = m.prompts[m.curField]
			m.ti.Focus()
			return m, cmd
		}
	}

	return m, cmd
}
