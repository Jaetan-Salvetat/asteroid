package settings

import (
	"asteroid/assets"
	"asteroid/core/geo"
	"asteroid/input"
	"asteroid/scene"
	"asteroid/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type SettingsScene struct {
	navigator  scene.Navigator
	backButton ui.TextButton
}

func NewScene(navigator scene.Navigator) *SettingsScene {
	backBtn := ui.NewTextButton("Go back", assets.Font(30), geo.Vector2{X: 30, Y: 30})

	return &SettingsScene{
		navigator:  navigator,
		backButton: backBtn,
	}
}

func (s *SettingsScene) Update(in input.Inputs) error {
	backBtnCLicked := s.backButton.Update(in)

	if backBtnCLicked {
		s.navigator.Pop()
	}

	return nil
}

func (s *SettingsScene) Draw(scene *ebiten.Image) {
	s.backButton.Draw(scene)
}

func (s *SettingsScene) OnEnter() {}

func (s *SettingsScene) OnExit() {}

func (s *SettingsScene) OnPause() {}

func (s *SettingsScene) OnResume() {}
