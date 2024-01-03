package title

import "github.com/charmbracelet/lipgloss"

type Option func(*Title)

func WithBackground(color lipgloss.Color) Option {
	return func(title *Title) {
		title.Background = color
	}
}

func WithForeground(color lipgloss.Color) Option {
	return func(title *Title) {
		title.Foreground = color
	}
}

func New(text string, options ...Option) Title {
	title := Title{
		Text:       text,
		Background: lipgloss.Color("#EB5E28"),
		Foreground: lipgloss.Color("#252422"),
	}

	for _, option := range options {
		option(&title)
	}

	return title
}

type Title struct {
	Text                   string
	Background, Foreground lipgloss.Color
}

func (t Title) String() string {
	return t.Text
}

func (t Title) Style() lipgloss.Style {
	return lipgloss.
		NewStyle().
		Background(t.Background).
		Foreground(t.Foreground).
		Bold(true).
		Padding(0, 1)
}

func (t Title) Render(parents ...lipgloss.Style) string {
	style := t.Style()

	for _, parent := range parents {
		style.Inherit(parent)
	}

	return style.Render(t.Text)
}
