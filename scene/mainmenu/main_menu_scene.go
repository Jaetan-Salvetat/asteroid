package mainmenu

import (
	"asteroid/core"
	"asteroid/scene"
	"asteroid/scene/settings"
	"asteroid/ui"
	"asteroid/ui/input"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"asteroid/assets"
)

type MainMenuScene struct {
	navigator      scene.Navigator
	settingsButton ui.TextButton
	background     ebiten.Image
}

func NewMainMenuScene(sm scene.Navigator) *MainMenuScene {
	settingsBtn := ui.NewTextButton("Go to settings", assets.Font(30), core.Vector2{X: 30, Y: 30})

	return &MainMenuScene{
		navigator:      sm,
		settingsButton: settingsBtn,
		background:     *assets.MenuBackground(),
	}
}

func (menu *MainMenuScene) Update(in input.Mouse) error {
	settingsPressed := menu.settingsButton.Update(in)

	if settingsPressed {
		menu.navigator.Push(settings.NewSettingsScene(menu.navigator))
	}

	return nil
}

func (menu *MainMenuScene) Draw(scene *ebiten.Image) {
	menu.settingsButton.Draw(scene)
	scene.DrawImage(&menu.background, &ebiten.DrawImageOptions{})
}

func (menu *MainMenuScene) OnEnter() {}

func (menu *MainMenuScene) OnExit() {}

func (menu *MainMenuScene) OnPause() {}

func (menu *MainMenuScene) OnResume() {}
