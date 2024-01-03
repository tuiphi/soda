package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
	"github.com/tuiphy/soda/title"
	"strings"
	"time"
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
	//return soda.NotifyWithDuration("Resized", time.Millisecond*300)
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
		case key.Matches(msg, s.keyMap.NextState):
			return soda.PushState(New(s.n + 1))
		case key.Matches(msg, s.keyMap.NextLayout):
			s.layout++
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

	fmt.Fprintf(&b, "State #%d", s.n)
	b.WriteString("\n\n")
	fmt.Fprintf(&b, "Available size\n%s", s.size)
	b.WriteString("\n\n")
	fmt.Fprintf(&b, "Focused: %t\n", s.focused)

	return b.String()
}
