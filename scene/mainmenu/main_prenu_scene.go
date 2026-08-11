package mainmenu

import (
	"asteroid/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenuScene struct {
	manager *scene.SceneManager
}

func NewMainMenuScene(sm *scene.SceneManager) *MainMenuScene {
	return &MainMenuScene{
		manager: sm,
	}
}

func (menu *MainMenuScene) Update() error {
	return nil
}

func (menu *MainMenuScene) Draw(scene *ebiten.Image) {}

func (menu *MainMenuScene) OnEnter() {}

func (menu *MainMenuScene) OnExit() {}

func (menu *MainMenuScene) OnPause() {}

func (menu *MainMenuScene) OnResume() {}