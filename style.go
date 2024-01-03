package soda

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Global,
	Header,
	Subtitle,
	Status,
	Spinner,
	Notification,
	Footer lipgloss.Style
}

func NewStyles() Styles {
	return Styles{
		Global: lipgloss.
			NewStyle().
			Padding(0, 1),
		Header: lipgloss.
			NewStyle().
			Padding(0, 0, 1, 0),
		Subtitle: lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}),
		Status: lipgloss.
			NewStyle().
			Padding(0, 1),
		Spinner: lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#8E8E8E", Dark: "#747373"}),
		Notification: lipgloss.
			NewStyle().
			Italic(true).
			Padding(0, 1).
			AlignHorizontal(lipgloss.Right),
		Footer: lipgloss.
			NewStyle().
			Padding(1, 1, 0, 0),
	}
}
