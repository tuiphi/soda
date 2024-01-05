package soda

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type stateWrapper struct {
	State

	SaveToHistory bool
}

type State interface {
	// Destroy is called when the state is destroyed
	Destroy()

	// Focused state will ask Model to ignore its KeyMap.
	// For example, pass "?" as tea.KeyMsg to the State instead of handling it by the Model (show help)
	Focused() bool

	// SetSize sets State's size.
	// The given Size is the Size of the terminal minus header and footer
	SetSize(size Size) tea.Cmd

	// Title of the State
	Title() string

	// Subtitle is shown below the Title
	Subtitle() string

	// Layout specifies a custom Layout for the state,
	// that will override the default one if the override is set to true
	Layout() (layout Layout, override bool)

	// Status is shown right to the Title
	Status() string

	// KeyMap of the State
	KeyMap() help.KeyMap

	// Init is the first function that will be called. It returns an optional
	// initial command. To not perform an initial command return nil.
	// This function is also called when the State is being popped out from the history
	Init(ctx context.Context) tea.Cmd

	// Update is called when a message is received. Use it to inspect messages
	// and, in response, update the model and/or send a command.
	Update(ctx context.Context, msg tea.Msg) tea.Cmd

	// View renders the State's UI, which is just a string. The view is
	// rendered after every Update.
	View() string
}
