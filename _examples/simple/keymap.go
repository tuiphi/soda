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
		NextState: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "next state"),
		),
	}
}

type KeyMap struct {
	ToggleSubtitle,
	SendNotification,
	NextState key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.ToggleSubtitle,
		k.SendNotification,
		k.NextState,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
