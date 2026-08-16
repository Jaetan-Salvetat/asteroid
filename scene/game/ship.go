package game

import (
	"asteroid/assets"
	"asteroid/config"
	"asteroid/core"
	"asteroid/core/geo"
	"asteroid/core/render"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const shootCooldown = 0.2

type Ship struct {
	Direction     geo.Vector2
	sprite        *ebiten.Image
	speed         float64
	rotation      float64
	Circle        geo.Circle
	shootCooldown core.Timer
}

func NewShip(position geo.Vector2) Ship {
	asset := assets.Ship.Cyan()
	return Ship{
		sprite:   asset,
		speed:    500,
		rotation: 0, //90 * math.Pi / 180,
		Circle: geo.Circle{
			Center: position,
			Radius: float64(asset.Bounds().Dx()) / 2,
		},
		shootCooldown: core.NewTimer(shootCooldown),
	}
}

func (s *Ship) Update(intent Intent) {
	s.shootCooldown.Update()
	s.move(intent.Movement)
	s.wrap()
	s.rotate(intent.Mouse.Cursor)
}

func (s *Ship) Draw(scene *ebiten.Image) {
	opt := render.ImageOptionsCentered(s.sprite.Bounds())
	opt.GeoM.Rotate(s.rotation)
	opt.GeoM.Translate(s.Circle.Center.X, s.Circle.Center.Y)
	scene.DrawImage(s.sprite, opt)
}

func (s *Ship) CanShoot(intent Intent) bool {
	return s.shootCooldown.IsExpired() && intent.Shooting
}

func (s *Ship) Shoot() Bullet {
	s.shootCooldown.Reset()

	return NewBullet(s.Circle.Center, s.Direction)
}

func (s *Ship) move(mouvement geo.Vector2) {
	s.Circle.Center = s.Circle.Center.Add(
		mouvement.Normalize().Scale(s.speed / float64(ebiten.TPS())),
	)
}

func (s *Ship) rotate(mouse geo.Vector2) {
	delta := s.Circle.Center.Delta(mouse)
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

	if s.Circle.Center.X+shipRayon <= leftLimit {
		s.Circle.Center.X = rightLimit + shipRayon
	} else if s.Circle.Center.X-shipRayon >= rightLimit {
		s.Circle.Center.X = leftLimit - shipRayon
	}

	if s.Circle.Center.Y+shipRayon <= topLimit {
		s.Circle.Center.Y = bottomLimit + shipRayon
	} else if s.Circle.Center.Y-shipRayon >= bottomLimit {
		s.Circle.Center.Y = topLimit - shipRayon
	}
}
