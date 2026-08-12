package game

import (
	"asteroid/assets"
	"asteroid/core/geo"
	"asteroid/scene"
	"asteroid/ui"
	"asteroid/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameScene struct {
	navigator  scene.Navigator
	backButton ui.TextButton
}

func NewGameScene(sm scene.Navigator) *GameScene {
	backBtn := ui.NewTextButton("Go back", assets.Font(30), geo.Vector2{X: 30, Y: 30})

	return &GameScene{
		navigator:  sm,
		backButton: backBtn,
	}
}

func (s *GameScene) Update(in input.Mouse) error {
	backBtnCLicked := s.backButton.Update(in)

	if backBtnCLicked {
		s.navigator.Pop()
	}

	return nil
}

func (s *GameScene) Draw(scene *ebiten.Image) {
	text.Draw(scene, "In game!", assets.Font(50), &text.DrawOptions{})
	s.backButton.Draw(scene)
}

func (s *GameScene) OnEnter() {}

func (s *GameScene) OnExit() {}

func (s *GameScene) OnPause() {}

func (s *GameScene) OnResume() {}
