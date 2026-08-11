package settings

import (
	"asteroid/assets"
	"asteroid/core"
	"asteroid/scene"
	"asteroid/ui"
	"asteroid/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type SettingsScene struct {
	navigator  scene.Navigator
	backButton ui.TextButton
}

func NewSettingsScene(sm scene.Navigator) *SettingsScene {
	backBtn := ui.NewTextButton("Go back", assets.Font(30), core.Vector2{X: 30, Y: 30})

	return &SettingsScene{
		navigator:  sm,
		backButton: backBtn,
	}
}

func (settings *SettingsScene) Update(in input.Mouse) error {
	backBtnCLicked := settings.backButton.Update(in)

	if backBtnCLicked {
		settings.navigator.Pop()
	}

	return nil
}

func (settings *SettingsScene) Draw(scene *ebiten.Image) {
	settings.backButton.Draw(scene)
}

func (settings *SettingsScene) OnEnter() {}

func (settings *SettingsScene) OnExit() {}

func (settings *SettingsScene) OnPause() {}

func (settings *SettingsScene) OnResume() {}
