package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/scene"
	"asteroid/scene/mainmenu"
	"asteroid/ui/input"
)

type Game struct {
	sceneManager *scene.SceneManager
	mouse        input.Mouse
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()

	g.mouse.Next(
		geo.Vector2{X: float64(x), Y: float64(y)},
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
	)

	return g.sceneManager.Update(g.mouse)
}

func (g *Game) Draw(scene *ebiten.Image) {
	g.sceneManager.Draw(scene)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var window = config.Window()
	return int(window.Width), int(window.Height)
}

func main() {
	var window = config.Window()
	var game = Game{}
	var sm = scene.NewSceneManager(game.mouse)

	game.sceneManager = sm
	sm.Push(mainmenu.NewMainMenuScene(sm))

	ebiten.SetWindowSize(int(window.Width/2), int(window.Height/2))
	ebiten.SetWindowTitle(config.AppName())

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
