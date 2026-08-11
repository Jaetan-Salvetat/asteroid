package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"asteroid/config"
	"asteroid/scene"
	"asteroid/scene/mainmenu"
)

type Game struct {
	sceneManager *scene.SceneManager
}

func (g *Game) Update() error {
	return g.sceneManager.Update()
}

func (g *Game) Draw(scene *ebiten.Image) {
	g.sceneManager.Draw(scene)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var window = config.Window()
	return window.Width, window.Height
}

func main() {
	var window = config.Window()
	var sm = scene.NewSceneManager()
	var game = Game{sceneManager: sm}

	sm.Push(mainmenu.NewMainMenuScene(sm))
	ebiten.SetWindowSize(window.Width, window.Height)
	ebiten.SetWindowTitle(config.AppName())
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
