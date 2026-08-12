package mainmenu

import (
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/scene"
	"asteroid/ui"
	"asteroid/ui/input"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"asteroid/assets"
)

type MainMenuScene struct {
	navigator  scene.Navigator
	background *ebiten.Image
	buttons    []ui.Button
}

func NewMainMenuScene(sm scene.Navigator) *MainMenuScene {
	centerX := float64(config.Window().Width) / 2

	return &MainMenuScene{
		navigator:  sm,
		background: assets.MenuBackground(),
		buttons: []ui.Button{
			ui.NewButton("PLAY", geo.SizeLarge, geo.Vector2{X: centerX, Y: 460}),
			ui.NewButton("SETTINGS", geo.SizeMedium, geo.Vector2{X: centerX, Y: 540}),
			ui.NewButton("QUIT", geo.SizeSmall, geo.Vector2{X: centerX, Y: 620}),
		},
	}
}

func (menu *MainMenuScene) Update(in input.Mouse) error {
	for i := range menu.buttons {
		menu.buttons[i].Update(in)
	}
	return nil
}

func (menu *MainMenuScene) Draw(scene *ebiten.Image) {
	scene.DrawImage(menu.background, &ebiten.DrawImageOptions{})

	for i := range menu.buttons {
		menu.buttons[i].Draw(scene)
	}
}

func (menu *MainMenuScene) OnEnter() {}

func (menu *MainMenuScene) OnExit() {}

func (menu *MainMenuScene) OnPause() {}

func (menu *MainMenuScene) OnResume() {}
