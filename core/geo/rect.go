package geo

type Rect struct{ X, Y, Width, Height float64 }

func (s Rect) Contains(p Vector2) bool {
	xEnd := s.X + s.Width
	yEnd := s.Y + s.Height

	return p.X >= s.X && p.X <= xEnd && p.Y >= s.Y && p.Y <= yEnd
}
