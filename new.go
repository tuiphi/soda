package soda

import (
	"context"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zyedidia/generic/stack"
)

func New(state State, options ...Option) *Model {
	ctx, ctxCancel := context.WithCancel(context.Background())

	model := &Model{
		styles: DefaultStyles(),
		state: stateWrapper{
			State:         state,
			SaveToHistory: true,
		},
		history: stack.New[stateWrapper](),
		onError: func(err error) tea.Cmd {
			const errorColor = lipgloss.Color("#ED4337")

			style := lipgloss.NewStyle().Bold(true).Foreground(errorColor)

			return Notify(style.Render(err.Error()))
		},
		spinner:                     spinner.Model{},
		showSpinner:                 false,
		size:                        Size{},
		keyMap:                      DefaultKeyMap(),
		help:                        help.New(),
		notificationDefaultDuration: time.Second * 3,
		notification:                _Notification{},
		defaultLayout:               NewLayout(),
		ctx:                         ctx,
		ctxCancel:                   ctxCancel,
		showHeader:                  true,
		showFooter:                  true,
	}

	for _, option := range options {
		option(model)
	}

	return model
}
