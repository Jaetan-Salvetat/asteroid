package mainmenu

import (
	"asteroid/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuBackgroud struct {
	backgroundImg *ebiten.Image
	shipImg *ebiten.Image
}

func NewMenuBackground() *MenuBackgroud {
	return &MenuBackgroud{
		backgroundImg: assets.MenuBackground(),
	}
}

func (s *MenuBackgroud) Update() error {
	return nil
}

func (s *MenuBackgroud) Draw(scene *ebiten.Image) {
	scene.DrawImage(s.backgroundImg, &ebiten.DrawImageOptions{})
}