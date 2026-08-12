package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"asteroid/core/geo"
	"asteroid/ui/input"
)

type TextButton struct {
	label   string
	face    *text.GoTextFace
	rect    geo.Rect
	state   input.Phase
	hovered bool
}

func NewTextButton(l string, f *text.GoTextFace, p geo.Vector2) TextButton {
	mWidth, mHeight := text.Measure(l, f, 0)

	return TextButton{label: l, face: f, rect: geo.Rect{X: p.X, Y: p.Y, Width: mWidth, Height: mHeight}}
}

func (s *TextButton) Height() float64 {
	_, height := text.Measure(s.label, s.face, 0)
	return height
}

func (s *TextButton) Update(in input.Mouse) bool {
	if s.rect.Contains(in.Cursor) {
		s.hovered = true
	} else {
		s.hovered = false
	}

	s.state = in.State

	return s.hovered && s.state == input.JustReleased
}

func (s *TextButton) Draw(scene *ebiten.Image) {
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(s.rect.X, s.rect.Y)

	text.Draw(scene, s.label, s.face, opts)
}
