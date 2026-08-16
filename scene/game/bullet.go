package game

import (
	"asteroid/assets"
	"asteroid/core"
	"asteroid/core/geo"
	"asteroid/core/render"

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

func (s *Bullet) Draw(scene *ebiten.Image) {
	opt := render.ImageOptionsCentered(s.sprite.Bounds())
	opt.GeoM.Translate(s.position.X, s.position.Y)
	scene.DrawImage(s.sprite, opt)
}

func (s *Bullet) IsAlive() bool {
	return !s.lifetime.IsExpired()
}

type Bullets []Bullet

func (s *Bullets) Update() {
	bullets := []Bullet{}

	for _, bullet := range *s {
		bullet.Update()

		if bullet.IsAlive() {
			bullets = append(bullets, bullet)
		}
	}

	*s = bullets
}

func (s *Bullets) Spawn(bullet Bullet) {
	*s = append(*s, bullet)
}

func (s *Bullets) Draw(scene *ebiten.Image) {
	for _, bullet := range *s {
		bullet.Draw(scene)
	}
}
