package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset"),
		),
		Increment: key.NewBinding(
			key.WithKeys("+"),
			key.WithHelp("+", "increment"),
		),
		Decrement: key.NewBinding(
			key.WithKeys("-"),
			key.WithHelp("-", "decrement"),
		),
	}
}

type KeyMap struct {
	Increment,
	Decrement,
	Reset key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Reset,
		k.Increment,
		k.Decrement,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
	}
}
