package ui

import (
	"asteroid/core/geo"
	"asteroid/input"
)

type ButtonState int

const (
	StateIdle ButtonState = iota
	StateHovered
	StateActive
	StateDisabled
)

func NewButtonStateFromMouse(mouse input.Mouse, btnRect geo.Rect) ButtonState {
	if btnRect.Contains(mouse.Cursor) {
		if mouse.IsDown() {
			return StateActive
		} else {
			return StateHovered
		}
	}

	return StateIdle
}
