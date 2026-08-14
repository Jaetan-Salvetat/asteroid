package game

import (
	"asteroid/assets"
	"asteroid/core/geo"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	sprite   *ebiten.Image
	speed    float64
	rotation    float64
	position geo.Vector2
}

func NewShip(position geo.Vector2) Ship {
	return Ship{
		sprite:   assets.ShipCyan(),
		speed:    20,
		rotation:    2,
		position: position,
	}
}
