package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
)

func run() error {
	model := soda.New(
		New(1),
		soda.WithMinSize(soda.Size{
			Width:  20,
			Height: 13,
		}),
	)

	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
