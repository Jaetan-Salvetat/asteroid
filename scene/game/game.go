package game

import (
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/core/render"
	"asteroid/input"
	"asteroid/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	navigator scene.Navigator
	ship      Ship
}

func NewScene(navigator scene.Navigator) *GameScene {
	shipX, shipY := config.Window().Width/2, config.Window().Height/2

	return &GameScene{
		navigator: navigator,
		ship:      NewShip(geo.Vector2{X: shipX, Y: shipY}),
	}
}

func (s *GameScene) Update(in input.Inputs) error {
	s.ship.Move(in)
	s.ship.Rotate(in.Mouse.Cursor)
	return nil
}

func (s *GameScene) Draw(scene *ebiten.Image) {
	opt := render.GeoMCentered(s.ship.sprite.Bounds())
	opt.GeoM.Rotate(s.ship.rotation)
	opt.GeoM.Translate(s.ship.position.X, s.ship.position.Y)
	scene.DrawImage(s.ship.sprite, opt)
}

func (s *GameScene) OnEnter() {}

func (s *GameScene) OnExit() {}

func (s *GameScene) OnPause() {}

func (s *GameScene) OnResume() {}
