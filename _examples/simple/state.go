package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
	"github.com/tuiphy/soda/title"
)

var _ soda.State = (*State)(nil)

type State struct {
	n      int
	size   soda.Size
	keyMap KeyMap

	layouts []soda.Layout
	layout  int

	focused bool

	showSubtitle bool
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	layout = s.layouts[s.layout%len(s.layouts)]

	return layout, true
}

func (s *State) Destroy() {
}

func (s *State) Focused() bool {
	return s.focused
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	s.size = size
	return nil
	// return soda.NotifyWithDuration("Resized", time.Millisecond*300)
}

func (s *State) Title() title.Title {
	return title.New("Simple")
}

func (s *State) Subtitle() string {
	if s.showSubtitle {
		return "Subtitle"
	}

	return ""
}

func (s *State) Status() string {
	return "Status"
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	return nil
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.ToggleSubtitle):
			s.showSubtitle = !s.showSubtitle
			return nil
		case key.Matches(msg, s.keyMap.SendNotification):
			return soda.NotifyWithDuration(time.Now().Format(time.StampMilli), time.Millisecond*800)
		case key.Matches(msg, s.keyMap.SendError):
			return func() tea.Msg {
				return errors.New("error: triggered manually")
			}
		case key.Matches(msg, s.keyMap.NextState):
			return soda.PushState(New(s.n + 1))
		case key.Matches(msg, s.keyMap.PrevLayout):
			if s.layout > 0 {
				s.layout--
			}

			return nil
		case key.Matches(msg, s.keyMap.NextLayout):
			if s.layout < len(s.layouts)-1 {
				s.layout++
			}

			return nil
		case key.Matches(msg, s.keyMap.ToggleFocus):
			s.focused = !s.focused
			return nil
		}
	}

	return nil
}

func (s *State) View() string {
	var b strings.Builder

	b.Grow(200)

	fmt.Fprintf(&b, "State #%d\n\n", s.n)
	fmt.Fprintf(&b, "Available size\n%s\n\n", s.size)

	b.WriteString("Focused: ")

	if s.focused {
		b.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#008000")).Render("true"))
		b.WriteString(" 🔒")
	} else {
		b.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ED4337")).Render("false"))
		b.WriteString(" 🔓")
	}
	b.WriteString("\n")

	b.WriteString(lipgloss.NewStyle().Faint(true).Render(`
Focused state will ask Model to ignore its KeyMap.
For example, pass "?" as tea.KeyMsg to the State instead of handling it by the Model (show help)
`))

	b.WriteString("\n")

	layout, _ := s.Layout()
	fmt.Fprintf(&b, "Layout\nHorizontal %.2f Vertical %.2f", layout.Horizontal, layout.Vertical)

	return b.String()
}
