package geo

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func GeoMCentered(size image.Rectangle) *ebiten.DrawImageOptions {
	width, height := size.Dx(), size.Dy()
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(width)/2, -float64(height)/2)

	return opt
}