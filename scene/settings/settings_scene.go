package settings

import (
	"asteroid/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type SettingsScene struct {
	manager *scene.SceneManager
}

func NewMainMenuScene(sm *scene.SceneManager) *SettingsScene {
	return &SettingsScene{
		manager: sm,
	}
}

func (menu *SettingsScene) Update() error {
	return nil
}

func (menu *SettingsScene) Draw(scene *ebiten.Image) {}

func (menu *SettingsScene) OnEnter() {}

func (menu *SettingsScene) OnExit() {}

func (menu *SettingsScene) OnPause() {}

func (menu *SettingsScene) OnResume() {}