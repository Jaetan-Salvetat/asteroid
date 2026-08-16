package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"asteroid/assets"
	"asteroid/core/geo"
	"asteroid/core/sound"
	"asteroid/input"
)

type Button struct {
	label         string
	face          *text.GoTextFace
	size          ButtonSize
	scale         float64
	rect          geo.Rect
	state         ButtonState
	previousState ButtonState
}

func NewButton(l string, s ButtonSize, p geo.Vector2) *Button {
	img := assets.Button.Idle()
	width := float64(img.Bounds().Dx()) * scale(s)
	height := float64(img.Bounds().Dy()) * scale(s)
	f := assets.Font(fontSize(s))

	return &Button{label: l, size: s, scale: scale(s), face: f, rect: geo.Rect{X: p.X, Y: p.Y, Width: width, Height: height}}
}

func (s *Button) Update(in input.Inputs) bool {
	s.previousState = s.state
	s.state = NewButtonStateFromMouse(in.Mouse, s.rect)

	s.playSound()

	return s.state == StateHovered && in.Mouse.State == input.JustReleased
}

func (s Button) Draw(scene *ebiten.Image) {
	imgOpts := &ebiten.DrawImageOptions{}
	imgOpts.GeoM.Scale(s.scale, s.scale)
	imgOpts.GeoM.Translate(s.rect.X, s.rect.Y)
	scene.DrawImage(s.image(), imgOpts)

	textOpts := &text.DrawOptions{}
	textOpts.LayoutOptions.PrimaryAlign = text.AlignCenter
	textOpts.LayoutOptions.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(s.rect.X+s.rect.Width/2, s.rect.Y+s.rect.Height/2)
	text.Draw(scene, s.label, s.face, textOpts)
}

func (s Button) Bounds() geo.Rect {
	return s.rect
}

func (s *Button) Place(vector geo.Vector2) {
	s.rect.X = vector.X
	s.rect.Y = vector.Y
}

func (s Button) image() *ebiten.Image {
	switch s.state {
	case StateActive:
		return assets.Button.Active()
	case StateHovered:
		return assets.Button.Hover()
	case StateDisabled:
		return assets.Button.Disabled()
	default:
		return assets.Button.Idle()
	}
}

func (s *Button) playSound() {
	if s.previousState == StateIdle && s.state == StateHovered {
		sound.Hover.Play()
	} else if s.previousState == StateHovered && s.state == StateActive {
		sound.Click.Play()
	}
}

func scale(size ButtonSize) float64 {
	switch size {
	case SizeSmall:
		return 1.2
	case SizeMedium:
		return 1.6
	default:
		return 2
	}
}

func fontSize(size ButtonSize) float64 {
	switch size {
	case SizeSmall:
		return 22
	case SizeMedium:
		return 29
	default:
		return 36
	}
}
