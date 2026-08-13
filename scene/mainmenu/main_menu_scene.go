package mainmenu

import (
	"asteroid/scene"
	"asteroid/ui/input"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"asteroid/assets"
)

type MainMenuScene struct {
	navigator  scene.Navigator
	background *ebiten.Image
	leftMenu   Menu
}

func NewMainMenuScene(navigator scene.Navigator) *MainMenuScene {
	return &MainMenuScene{
		navigator:  navigator,
		background: assets.MenuBackground(),
		leftMenu:   *NewMenu(navigator),
	}
}

func (s *MainMenuScene) Update(in input.Mouse) error {
	error := s.leftMenu.Update(in)
	if error != nil {
		return error
	}

	return nil
}

func (s *MainMenuScene) Draw(scene *ebiten.Image) {
	scene.DrawImage(s.background, &ebiten.DrawImageOptions{})
	s.leftMenu.Draw(scene)
}

func (s *MainMenuScene) OnEnter() {}

func (s *MainMenuScene) OnExit() {}

func (s *MainMenuScene) OnPause() {}

func (s *MainMenuScene) OnResume() {}
