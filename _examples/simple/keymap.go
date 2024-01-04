package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func NewKeyMap() KeyMap {
	return KeyMap{
		ToggleSubtitle: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle subtitle"),
		),
		SendNotification: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "send notification"),
		),
		SendError: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "send error"),
		),
		NextState: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "next state"),
		),
		PrevLayout: key.NewBinding(
			key.WithKeys("L"),
			key.WithHelp("L", "prev layout"),
		),
		NextLayout: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "next layout"),
		),
		ToggleFocus: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "toggle focus"),
		),
	}
}

type KeyMap struct {
	ToggleSubtitle,
	SendNotification,
	SendError,
	NextState,
	NextLayout,
	PrevLayout,
	ToggleFocus key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
		{
			k.PrevLayout,
			k.NextLayout,
			k.NextState,
		},
		{
			k.ToggleFocus,
			k.ToggleSubtitle,
		},
		{
			k.SendNotification,
			k.SendError,
		},
	}
}
