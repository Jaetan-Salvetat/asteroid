package input

import "asteroid/core/geo"

type ButtonState uint8

const (
	StateIdle ButtonState = iota
	StateHovered
	StateActive
	StateDisabled
)

func NewButtonStateFromMouse(mouse Mouse, btnRect geo.Rect) ButtonState {
	if btnRect.Contains(mouse.Cursor) {
		if mouse.IsDown() {
			return StateActive
		} else {
			return StateHovered
		}
	}
	
	return StateIdle
}