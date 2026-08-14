package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/input"
	"asteroid/scene"
	"asteroid/scene/mainmenu"
)

type Game struct {
	sceneManager *scene.SceneManager
	inputs       input.Inputs
}

func NewGame() Game {
	return Game{inputs: input.NewInputs()}
}

func (g *Game) Update() error {
	pressedKeys := []ebiten.Key{}
	x, y := ebiten.CursorPosition()

	g.inputs.Mouse.Next(
		geo.Vector2{X: float64(x), Y: float64(y)},
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
	)

	pressedKeys = inpututil.AppendPressedKeys(pressedKeys[:0])
	g.inputs.Keyboard.Next(pressedKeys)

	return g.sceneManager.Update(g.inputs)
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
	var game = NewGame()
	game.sceneManager = scene.NewManager()

	game.sceneManager.Push(mainmenu.NewScene(game.sceneManager))

	ebiten.SetWindowSize(int(window.Width/2), int(window.Height/2))
	ebiten.SetWindowTitle(config.AppName)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
