package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tuiphy/soda"
	"log"
)

func run() error {
	state := New()
	model := soda.New(state)
	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
