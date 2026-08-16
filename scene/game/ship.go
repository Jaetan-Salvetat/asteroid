package game

import (
	"asteroid/assets"
	"asteroid/config"
	"asteroid/core"
	"asteroid/core/geo"
	"asteroid/input"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const shootCooldown = 0.2

type Ship struct {
	Direction     geo.Vector2
	sprite        *ebiten.Image
	speed         float64
	rotation      float64
	position      geo.Vector2
	shootCooldown core.Timer
}

func NewShip(position geo.Vector2) Ship {
	return Ship{
		sprite:        assets.ShipCyan(),
		speed:         500,
		rotation:      0, //90 * math.Pi / 180,
		position:      position,
		shootCooldown: core.NewTimer(shootCooldown),
	}
}

func (s *Ship) Update(in input.Inputs) {
	s.shootCooldown.Update()
	s.move(in)
	s.wrap()
	s.rotate(in.Mouse.Cursor)
}

func (s *Ship) CanShoot() bool {
	return s.shootCooldown.IsExpired()
}

func (s *Ship) Shoot() Bullet {
	s.shootCooldown.Reset()

	return NewBullet(s.position, s.Direction)
}

func (s *Ship) move(in input.Inputs) {
	dx, dw := 0.0, 0.0

	if in.Keyboard.Pressed(ebiten.KeyW) {
		dw -= 1
	}
	if in.Keyboard.Pressed(ebiten.KeyS) {
		dw += 1
	}
	if in.Keyboard.Pressed(ebiten.KeyA) {
		dx -= 1
	}
	if in.Keyboard.Pressed(ebiten.KeyD) {
		dx += 1
	}

	s.position = s.position.Add(geo.Vector2{X: dx, Y: dw}.Normalize().Scale(s.speed / float64(ebiten.TPS())))
}

func (s *Ship) rotate(mouse geo.Vector2) {
	delta := s.position.Delta(mouse)
	rotation := math.Atan2(delta.Y, delta.X)
	s.rotation = rotation
	s.Direction = geo.Vector2{
		X: math.Cos(rotation),
		Y: math.Sin(rotation),
	}
}

func (s *Ship) wrap() {
	window := config.Window()
	shipRayon := float64(s.sprite.Bounds().Dx()) / 2
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
