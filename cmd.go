package soda

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

// Back to the previous state
func Back() tea.Msg {
	return backMsg{Steps: 1}
}

// BackN traverses the history N states back
func BackN(n int) tea.Cmd {
	if n < 0 {
		panic("n < 0")
	}

	return func() tea.Msg {
		return backMsg{Steps: n}
	}
}

// BackToRoot traverses to the first (initial) State in the history
func BackToRoot() tea.Msg {
	return backToRootMsg{}
}

// PushState will push a new State
func PushState(state State) tea.Cmd {
	return func() tea.Msg {
		return pushStateMsg{State: stateWrapper{
			State:         state,
			SaveToHistory: true,
		}}
	}
}

// PushTempState will push a new State that won't be saved into history
func PushTempState(state State) tea.Cmd {
	return func() tea.Msg {
		return pushStateMsg{State: stateWrapper{
			State:         state,
			SaveToHistory: false,
		}}
	}
}

// Notify sends a notification with the default time.Duration
func Notify(message string) tea.Cmd {
	return func() tea.Msg {
		return notificationMsg{Message: message}
	}
}

// NotifyWithDuration sends a notification with the given time.Duration ignoring the default
func NotifyWithDuration(message string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return notificationWithDurationMsg{
			notificationMsg: notificationMsg{
				Message: message,
			},
			Duration: duration,
		}
	}
}
