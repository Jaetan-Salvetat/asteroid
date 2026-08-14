package ui

import (
	"asteroid/core/geo"
)

type Column struct {
	rect              geo.Rect
	alignment         Alignment
	verticalAlignment Alignment
	gap               float64
	children          []Widget
}

func NewColumn(rect geo.Rect, alignment Alignment, verticalAlignment Alignment, gap float64, chidren ...Widget) *Column {
	return &Column{
		rect:              rect,
		alignment:         alignment,
		verticalAlignment: verticalAlignment,
		gap:               gap,
		children:          chidren,
	}
}

func (s *Column) Place() {
	y := s.rect.Y

	for _, child := range s.children {
		b := child.Bounds()
		child.Place(geo.Vector2{X: s.rect.X + s.offsetX(b.Width), Y: y + s.offsetY()})
		y += b.Height + s.gap
	}
}

func (s *Column) offsetX(width float64) float64 {
	switch s.alignment {
	case AlignCenter:
		return (s.rect.Width - width) / 2
	case AlignEnd:
		return s.rect.Width - width
	default:
		return 0
	}
}

func (s *Column) offsetY() float64 {
	switch s.verticalAlignment {
	case AlignCenter:
		return (s.rect.Height - s.contentHeight()) / 2
	case AlignEnd:
		return s.rect.Height - s.contentHeight()
	default:
		return 0
	}
}

func (s *Column) contentHeight() float64 {
	height := float64(s.gap)

	for i, child := range s.children {
		height += child.Bounds().Height

		if i > 0 {
			height += s.gap
		}
	}

	return height
}
