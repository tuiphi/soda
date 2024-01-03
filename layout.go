package soda

import "github.com/charmbracelet/lipgloss"

type Layout struct {
	Horizontal, Vertical lipgloss.Position
}

func NewLayout() Layout {
	return Layout{
		Horizontal: lipgloss.Left,
		Vertical:   lipgloss.Top,
	}
}
