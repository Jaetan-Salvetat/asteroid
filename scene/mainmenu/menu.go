package mainmenu

import (
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/scene"
	"asteroid/scene/game"
	"asteroid/scene/settings"
	"asteroid/ui"
	"asteroid/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	navigator      scene.Navigator
	startButton    *ui.Button
	settingsButton *ui.Button
	quitButton     *ui.Button
}

func NewMenu(navigator scene.Navigator, width float64) *Menu {

	menu := &Menu{
		navigator:      navigator,
		startButton:    ui.NewButton("PLAY", geo.SizeLarge, geo.Vector2{X: 0, Y: 0}),
		settingsButton: ui.NewButton("SETTINGS", geo.SizeLarge, geo.Vector2{X: 0, Y: 0}),
		quitButton:     ui.NewButton("QUIT", geo.SizeLarge, geo.Vector2{X: 0, Y: 0}),
	}

	column := ui.NewColumn(
		geo.Rect{X: 0, Y: 0, Width: float64(width), Height: float64(config.Window().Height)},
		input.AlignCenter, input.AlignCenter, 30,
		menu.startButton, menu.settingsButton, menu.quitButton,
	)
	column.Place()

	return menu
}

func (s *Menu) Update(in input.Inputs) error {
	startClicked := s.startButton.Update(in)
	settingsClicked := s.settingsButton.Update(in)
	quitClicked := s.quitButton.Update(in)

	if startClicked {
		s.navigator.Push(game.NewGameScene(s.navigator))
	}
	if settingsClicked {
		s.navigator.Push(settings.NewSettingsScene(s.navigator))
	}
	if quitClicked {
		return ebiten.Termination
	}

	return nil
}

func (s *Menu) Draw(scene *ebiten.Image) {
	s.startButton.Draw(scene)
	s.settingsButton.Draw(scene)
	s.quitButton.Draw(scene)
}
