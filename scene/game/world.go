package game

import (
	"asteroid/config"
	"asteroid/core/geo"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	ship      Ship
	bullets   Bullets
	asteroids Asteroids
}

func NewWorld() World {
	windowCenter := geo.Vector2{
		X: config.Window().Width / 2,
		Y: config.Window().Height / 2,
	}

	return World{ship: NewShip(windowCenter)}
}

func (s *World) Update(intent Intent) Frame {
	s.ship.Update(intent)
	s.bullets.Update()

	if s.ship.CanShoot(intent) {
		s.bullets.Spawn(s.ship.Shoot())
	}

	return s.handleCollisions()
}

func (s *World) Draw(scene *ebiten.Image) {
	s.bullets.Draw(scene)
	s.ship.Draw(scene)
	s.asteroids.Draw(scene)
}

func (s *World) IsCleared() bool {
	return len(s.asteroids) <= 0
}

func (s *World) SpawnWave(ship geo.Circle) {
	s.asteroids.SpawnWaved(ship)
}

func (s *World) handleCollisions() Frame {
	return Frame{}
}
