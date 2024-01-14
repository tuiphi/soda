package main

import (
	"context"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/tuiphy/soda"
)

var _ soda.State = (*State)(nil)

type State struct {
	counter int64
	keyMap  KeyMap
}

func (s *State) Destroy() {
}

func (s *State) Focused() bool {
	return false
}

func (s *State) SetSize(size soda.Size) tea.Cmd {
	return nil
}

func (s *State) Title() string {
	return "Counter"
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Layout() (layout soda.Layout, override bool) {
	return soda.Layout{
		Horizontal: lipgloss.Center,
		Vertical:   lipgloss.Center,
	}, true
}

func (s *State) Status() string {
	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(ctx context.Context) tea.Cmd {
	s.keyMap.Decrement.SetEnabled(false)
	return nil
}

func (s *State) Update(ctx context.Context, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Reset):
			s.setValue(0)
			return nil
		case key.Matches(msg, s.keyMap.Increment):
			s.increment()
			return nil
		case key.Matches(msg, s.keyMap.Decrement):
			s.decrement()
			return nil
		}
	}

	return nil
}

func (s *State) increment() {
	s.setValue(s.counter + 1)
}

func (s *State) decrement() {
	s.setValue(s.counter - 1)
}

func (s *State) setValue(value int64) {
	s.counter = value

	if s.counter == 0 {
		s.keyMap.Decrement.SetEnabled(false)
	} else {
		s.keyMap.Decrement.SetEnabled(true)
	}
}

func (s *State) View(soda.Layout) string {
	t := table.
		New().
		Border(lipgloss.NormalBorder()).
		Headers("Base", "Value").
		Rows(
			[]string{"Binary", strconv.FormatInt(s.counter, 2)},
			[]string{"Octal", strconv.FormatInt(s.counter, 8)},
			[]string{"Decimal", strconv.FormatInt(s.counter, 10)},
			[]string{"Hexadecimal", strconv.FormatInt(s.counter, 16)},
		)

	return t.String()
}
