package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"asteroid/core"
	"asteroid/ui/input"
)

type TextButton struct {
	label   string
	face    *text.GoTextFace
	rect    core.Rect
	state   input.Phase
	hovered bool
}

func NewTextButton(l string, f *text.GoTextFace, p core.Vector2) TextButton {
	mWidth, mHeight := text.Measure(l, f, 0)

	return TextButton{label: l, face: f, rect: core.Rect{X: p.X, Y: p.Y, Width: mWidth, Height: mHeight}}
}

func (btn *TextButton) Update(in input.Mouse) bool {
	if btn.rect.Contains(in.Cursor) {
		btn.hovered = true
	} else {
		btn.hovered = false
	}

	btn.state = in.State

	return btn.hovered && btn.state == input.JustReleased
}

func (btn *TextButton) Draw(scene *ebiten.Image) {
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(btn.rect.X, btn.rect.Y)

	text.Draw(scene, btn.label, btn.face, opts)
}
