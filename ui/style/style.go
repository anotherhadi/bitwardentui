package style

import "github.com/charmbracelet/lipgloss"

var (

	// Colors

	Foreground = lipgloss.Color("#FCFCFC")
	Background = lipgloss.Color("#101010")
	Accent     = lipgloss.Color("39")
	Muted      = lipgloss.Color("244")

	// Styles

	Text        = lipgloss.NewStyle().Foreground(Foreground)
	TextMuted   = lipgloss.NewStyle().Foreground(Muted)
	TextFocused = lipgloss.NewStyle().Foreground(Accent)

	ButtonFocused = lipgloss.NewStyle().Foreground(Accent).Padding(0, 1).Border(lipgloss.RoundedBorder()).BorderForeground(Accent)
	ButtonMuted   = lipgloss.NewStyle().Foreground(Muted).Padding(0, 1).Border(lipgloss.RoundedBorder()).BorderForeground(Muted)

	Box = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(Muted)

	CursorStyle = TextFocused
)
