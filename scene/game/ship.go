package game

import (
	"asteroid/assets"
	"asteroid/core/geo"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	sprite   *ebiten.Image
	speed    float64
	angle    float64
	position geo.Vector2
}

func NewShip(position geo.Vector2) Ship {
	return Ship{
		sprite:   assets.ShipCyan(),
		speed:    20,
		angle:    2,
		position: position,
	}
}
