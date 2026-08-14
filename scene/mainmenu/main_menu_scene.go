package mainmenu

import (
	"asteroid/assets"
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/scene"
	"asteroid/ui/input"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenuScene struct {
	navigator  scene.Navigator
	background *ebiten.Image
	ship       *MenuShip
	menu       *Menu
}

func NewMainMenuScene(navigator scene.Navigator) *MainMenuScene {
	menuWidth := config.Window().Width / 2
	shipX := config.Window().Width * 0.65
	shipY := config.Window().Height / 2

	return &MainMenuScene{
		navigator:  navigator,
		background: assets.MenuBackground(),
		ship:       NewMenuShip(geo.Vector2{X: shipX, Y: shipY}),
		menu:       NewMenu(navigator, menuWidth),
	}
}

func (s *MainMenuScene) Update(in input.Inputs) error {
	error := s.menu.Update(in)
	if error != nil {
		return error
	}

	error = s.ship.Update()
	if error != nil {
		return error
	}

	return nil
}

func (s *MainMenuScene) Draw(scene *ebiten.Image) {
	scene.DrawImage(s.background, &ebiten.DrawImageOptions{})
	s.ship.Draw(scene)
	s.menu.Draw(scene)
}

func (s *MainMenuScene) OnEnter() {}

func (s *MainMenuScene) OnExit() {}

func (s *MainMenuScene) OnPause() {}

func (s *MainMenuScene) OnResume() {}
