package game

import (
	"asteroid/assets"
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/input"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	sprite   *ebiten.Image
	speed    float64
	rotation float64
	position geo.Vector2
}

func NewShip(position geo.Vector2) Ship {
	return Ship{
		sprite:   assets.ShipCyan(),
		speed:    500,
		rotation: 0, //90 * math.Pi / 180,
		position: position,
	}
}

func (s *Ship) Move(in input.Inputs) {
	if in.Keyboard.Pressed(ebiten.KeyW) {
		s.position = s.position.Add(0, -(s.speed / float64(ebiten.TPS())))
	} else if in.Keyboard.Pressed(ebiten.KeyS) {
		s.position = s.position.Add(0, s.speed/float64(ebiten.TPS()))
	}

	if in.Keyboard.Pressed(ebiten.KeyA) {
		s.position = s.position.Add(-(s.speed / float64(ebiten.TPS())), 0)
	} else if in.Keyboard.Pressed(ebiten.KeyD) {
		s.position = s.position.Add(s.speed/float64(ebiten.TPS()), 0)
	}
}

func (s *Ship) Rotate(mouse geo.Vector2) {
	delta := s.position.DirectionTo(mouse)
	rotation := math.Atan2(delta.Y, delta.X)
	s.rotation = rotation
}

func (s *Ship) Wrap() {
	window := config.Window()
	shipRayon := float64(s.sprite.Bounds().Dx() / 2)
	leftLimit := 0.0
	rightLimit := window.Width
	topLimit := 0.0
	bottomLimit := window.Height

	if s.position.X+shipRayon <= leftLimit {
		s.position.X = rightLimit + shipRayon
	} else if s.position.X-shipRayon >= rightLimit {
		s.position.X = leftLimit - shipRayon
	}

	if s.position.Y+shipRayon <= topLimit {
		s.position.Y = bottomLimit + shipRayon
	} else if s.position.Y-shipRayon >= bottomLimit {
		s.position.Y = topLimit - shipRayon
	}
}
