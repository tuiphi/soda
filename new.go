package soda

import (
	"context"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zyedidia/generic/stack"
	"time"
)

func New(state State, options ...Option) *Model {
	ctx, ctxCancel := context.WithCancel(context.Background())

	model := &Model{
		styles: NewStyles(),
		state: stateWrapper{
			State:         state,
			SaveToHistory: true,
		},
		history: stack.New[stateWrapper](),
		onError: func(err error) tea.Cmd {
			return Notify(err.Error())
		},
		spinner:                     spinner.Model{},
		showSpinner:                 false,
		size:                        Size{},
		keyMap:                      NewKeyMap(),
		help:                        help.New(),
		notificationDefaultDuration: time.Second * 3,
		notification:                notification{},
		ctx:                         ctx,
		ctxCancel:                   ctxCancel,
	}

	for _, option := range options {
		option(model)
	}

	return model
}
