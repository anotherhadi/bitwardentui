package login

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anotherhadi/bitwarden_tui/ui/style"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	values      = []string{}
	send   bool = false
)

type model struct {
	focusIndex    int
	inputsLabel   []string
	inputs        []textinput.Model
	width, height int
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CharLimit = 32
		values = append(values, t.Value())

		switch i {
		case 0:
			m.inputsLabel = append(m.inputsLabel, "Server URL")
			t.Placeholder = "https://bitwarden.com"
			t.Focus()
			t.PromptStyle = style.TextFocused
			t.TextStyle = style.TextFocused
			t.CharLimit = 124
		case 1:
			m.inputsLabel = append(m.inputsLabel, "Email")
			t.Placeholder = "example@example.com"
			t.CharLimit = 64
			t.PromptStyle = style.Text
			t.TextStyle = style.Text
		case 2:
			m.inputsLabel = append(m.inputsLabel, "Password")
			t.Placeholder = "password123"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
			t.PromptStyle = style.Text
			t.TextStyle = style.Text
		}

		m.inputs[i] = t
	}

	m.width = 0
	m.height = 0

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// update size
		m.width, m.height = msg.Width, msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				send = true
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = style.TextFocused
					m.inputs[i].TextStyle = style.TextFocused
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = style.Text
				m.inputs[i].TextStyle = style.Text
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		values[i] = m.inputs[i].Value()
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(" _             _\n")
	b.WriteString("| | ___   __ _(_)_ __\n")
	b.WriteString("| |/ _ \\ / _` | | '_ \\\n")
	b.WriteString("| | (_) | (_| | | | | |\n")
	b.WriteString("|_|\\___/ \\__, |_|_| |_|\n")
	b.WriteString("         |___/\n\n\n")

	for i := range m.inputs {
		// label
		b.WriteString(m.inputsLabel[i] + "\n")
		// input
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteString("\n\n")
		}
	}

	button := style.ButtonMuted.Render("Submit")
	if m.focusIndex == len(m.inputs) {
		button = style.ButtonFocused.Render("Submit")
	}
	fmt.Fprintf(&b, "\n\n%s", button)

	box := style.Box.Width(60)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box.Render(b.String()))
}

func Login() (serverUrl, email, password string, err error) {
	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		return "", "", "", err
	}
	if send == false {
		return "", "", "", errors.New("login canceled")
	}
	return values[0], values[1], values[2], nil
}
