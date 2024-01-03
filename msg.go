package soda

import (
	"github.com/charmbracelet/bubbles/spinner"
	"time"
)

type (
	_NotificationMsg struct {
		Message string
	}

	_NotificationWithDurationMsg struct {
		_NotificationMsg

		Duration time.Duration
	}

	_NotificationTimeoutMsg struct{}

	_BackMsg struct {
		Steps int
	}

	_BackToRootMsg struct{}

	_PushStateMsg struct {
		State stateWrapper
	}

	_SpinnerTickMsg spinner.TickMsg
)
