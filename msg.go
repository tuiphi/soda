package soda

import (
	"github.com/charmbracelet/bubbles/spinner"
	"time"
)

type (
	notificationMsg struct {
		Message string
	}

	notificationWithDurationMsg struct {
		notificationMsg

		Duration time.Duration
	}

	notificationTimeoutMsg struct{}

	backMsg struct {
		Steps int
	}

	backToRootMsg struct{}

	pushStateMsg struct {
		State stateWrapper
	}

	spinnerTickMsg spinner.TickMsg
)
