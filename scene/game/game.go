package game

import (
	"asteroid/assets"
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/core/render"
	"asteroid/input"
	"asteroid/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	navigator  scene.Navigator
	background *ebiten.Image
	ship       Ship
	bullets    []Bullet
}

func NewScene(navigator scene.Navigator) *GameScene {
	shipX, shipY := config.Window().Width/2, config.Window().Height/2

	return &GameScene{
		navigator:  navigator,
		background: assets.Background(),
		ship:       NewShip(geo.Vector2{X: shipX, Y: shipY}),
		bullets:    []Bullet{},
	}
}

func (s *GameScene) Update(in input.Inputs) error {
	s.ship.Update(in)
	s.updateBullets()

	if s.ship.CanShoot() && in.Mouse.IsDown() {
		s.bullets = append(s.bullets, s.ship.Shoot())
	}

	return nil
}

func (s *GameScene) Draw(scene *ebiten.Image) {
	scene.DrawImage(s.background, &ebiten.DrawImageOptions{})
	s.drawBullets(scene)
	s.drawShip(scene)
}

func (s *GameScene) OnEnter() {}

func (s *GameScene) OnExit() {}

func (s *GameScene) OnPause() {}

func (s *GameScene) OnResume() {}

func (s *GameScene) drawShip(scene *ebiten.Image) {
	opt := render.ImageOptionsCentered(s.ship.sprite.Bounds())
	opt.GeoM.Rotate(s.ship.rotation)
	opt.GeoM.Translate(s.ship.position.X, s.ship.position.Y)
	scene.DrawImage(s.ship.sprite, opt)
}

func (s *GameScene) drawBullets(scene *ebiten.Image) {
	for _, bullet := range s.bullets {
		opt := render.ImageOptionsCentered(bullet.sprite.Bounds())
		opt.GeoM.Translate(bullet.position.X, bullet.position.Y)
		scene.DrawImage(bullet.sprite, opt)
	}
}

func (s *GameScene) updateBullets() {
	bullets := []Bullet{}

	for i := range s.bullets {
		s.bullets[i].Update()

		if s.bullets[i].IsAlive() {
			bullets = append(bullets, s.bullets[i])
		}
	}

	s.bullets = bullets
}
