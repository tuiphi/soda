package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/tuiphy/soda"
)

func New(n int) *State {
	var layouts []soda.Layout

	positions := []lipgloss.Position{
		0,
		0.5,
		1,
	}

	for _, x := range positions {
		for _, y := range positions {
			layouts = append(layouts, soda.Layout{
				Horizontal: x,
				Vertical:   y,
			})
		}
	}

	return &State{
		n:       n,
		keyMap:  NewKeyMap(),
		layouts: layouts,
	}
}
