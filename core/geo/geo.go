package geo

type Vector2 struct{ X, Y float64 }
type Rect struct{ X, Y, Width, Height float64 }

func (s Rect) Contains(p Vector2) bool {
	xEnd := s.X + s.Width
	yEnd := s.Y + s.Height

	return p.X >= s.X && p.X <= xEnd && p.Y >= s.Y && p.Y <= yEnd
}

func (s Vector2) Add(x, y float64) Vector2 {
	return Vector2{X: s.X + x, Y: s.Y + y}
}

func (s Vector2) DirectionTo(to Vector2) Vector2 {
	return Vector2{
		X: to.X - s.X,
		Y: to.Y - s.Y,
	}
}
