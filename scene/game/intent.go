package game

import (
	"asteroid/core/geo"
	"asteroid/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Intent struct {
	Movement geo.Vector2
	Mouse    input.Mouse
	Shooting bool
}

func NewIntentFrom(inputs input.Inputs) Intent {
	movement := geo.Vector2{}

	if inputs.Keyboard.Pressed(ebiten.KeyW) {
		movement = movement.AddY(-1)
	}
	if inputs.Keyboard.Pressed(ebiten.KeyS) {
		movement = movement.AddY(1)
	}
	if inputs.Keyboard.Pressed(ebiten.KeyA) {
		movement = movement.AddX(-1)
	}
	if inputs.Keyboard.Pressed(ebiten.KeyD) {
		movement = movement.AddX(1)
	}

	return Intent{
		Movement: movement,
		Mouse:    inputs.Mouse,
		Shooting: inputs.Mouse.IsDown(),
	}
}
