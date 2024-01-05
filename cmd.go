package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Back to the previous state
func Back() tea.Msg {
	return _BackMsg{Steps: 1}
}

// BackN traverses the history N states back
func BackN(n int) tea.Cmd {
	if n < 0 {
		panic("n < 0")
	}

	return func() tea.Msg {
		return _BackMsg{Steps: n}
	}
}

// BackToRoot traverses to the first (initial) State in the history
func BackToRoot() tea.Msg {
	return _BackToRootMsg{}
}

// PushState will push a new State
func PushState(state State) tea.Cmd {
	return func() tea.Msg {
		return _PushStateMsg{State: stateWrapper{
			State:         state,
			SaveToHistory: true,
		}}
	}
}

// PushTempState will push a new State that won't be saved into history
func PushTempState(state State) tea.Cmd {
	return func() tea.Msg {
		return _PushStateMsg{State: stateWrapper{
			State:         state,
			SaveToHistory: false,
		}}
	}
}

// Notify sends a notification with the default time.Duration
func Notify(message string) tea.Cmd {
	return func() tea.Msg {
		return _NotificationMsg{Message: message}
	}
}

// NotifyWithDuration sends a notification with the given time.Duration ignoring the default
func NotifyWithDuration(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return _NotificationWithDurationMsg{
			_NotificationMsg: _NotificationMsg{
				Message: message,
			},
			Duration: duration,
		}
	}
}

func SendError(err error) tea.Cmd {
	return func() tea.Msg {
		return err
	}
}

func Wrap(supplier func() tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		cmd := supplier()

		if cmd == nil {
			return nil
		}

		return cmd()
	}
}
