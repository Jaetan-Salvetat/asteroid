package ui

import (
	_ "image/png"

	_ "golang.org/x/image/webp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"asteroid/assets"
	"asteroid/core/geo"
	"asteroid/core/sound"
	"asteroid/ui/input"
)

type Button struct {
	label         string
	face          *text.GoTextFace
	size          geo.ButtonSize
	scale         float64
	rect          geo.Rect
	mouse         input.Mouse
	state         input.ButtonState
	previousState input.ButtonState
}

func NewButton(l string, s geo.ButtonSize, p geo.Vector2) *Button {
	img := assets.ButtonIdle()
	width := float64(img.Bounds().Dx()) * scale(s)
	height := float64(img.Bounds().Dy()) * scale(s)
	f := assets.Font(fontSize(s))

	return &Button{label: l, size: s, scale: scale(s), face: f, rect: geo.Rect{X: p.X, Y: p.Y, Width: width, Height: height}}
}

func (s *Button) Height() float64 {
	_, height := text.Measure(s.label, s.face, 0)
	return height
}

func (s *Button) Update(in input.Mouse) bool {
	s.previousState = s.state
	s.state = input.NewButtonStateFromMouse(in, s.rect)
	s.mouse = in

	s.playSound()

	return s.state == input.StateHovered && in.State == input.JustReleased
}

func (s *Button) Draw(scene *ebiten.Image) {
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

func (s *Button) Bounds() geo.Rect {
	return s.rect
}

func (s *Button) Place(vector geo.Vector2) {
	s.rect.X = vector.X
	s.rect.Y = vector.Y
}

func (s *Button) image() *ebiten.Image {
	switch s.state {
	case input.StateActive:
		return assets.ButtonActive()
	case input.StateHovered:
		return assets.ButtonHovered()
	case input.StateDisabled:
		return assets.ButtonDisabled()
	default:
		return assets.ButtonIdle()
	}
}

func (s *Button) playSound() {
	if s.previousState == input.StateIdle && s.state == input.StateHovered {
		sound.Hover.Play()
	} else if s.previousState == input.StateHovered && s.state == input.StateActive {
		sound.Click.Play()
	}
}

func scale(size geo.ButtonSize) float64 {
	switch size {
	case geo.SizeSmall:
		return 1.2
	case geo.SizeMedium:
		return 1.6
	default:
		return 2
	}
}

func fontSize(size geo.ButtonSize) float64 {
	switch size {
	case geo.SizeSmall:
		return 22
	case geo.SizeMedium:
		return 29
	default:
		return 36
	}
}
