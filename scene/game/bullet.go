package game

import (
	"asteroid/assets"
	"asteroid/core"
	"asteroid/core/geo"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	sprite    *ebiten.Image
	position  geo.Vector2
	direction geo.Vector2
	speed     float64
	damage    int
	lifetime  core.Timer
}

func NewBullet(position geo.Vector2, direction geo.Vector2) Bullet {
	return Bullet{
		sprite:    assets.UfoShoot(),
		position:  position,
		direction: direction,
		speed:     1000,
		damage:    1,
		lifetime:  core.NewTimer(1),
	}
}

func (s *Bullet) Update() {
	s.lifetime.Update()
	s.position = s.position.Add(s.direction.Normalize().Scale(s.speed / float64(ebiten.TPS())))
}

func (s *Bullet) IsAlive() bool {
	return !s.lifetime.IsExpired()
}
