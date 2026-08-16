package geo

import "math"

type Vector2 struct{ X, Y float64 }

func (s Vector2) AddX(x float64) Vector2 {
	return Vector2{X: s.X + x, Y: s.Y}
}

func (s Vector2) AddY(y float64) Vector2 {
	return Vector2{X: s.X, Y: s.Y + y}
}

func (s Vector2) Add(v Vector2) Vector2 {
	return Vector2{X: s.X + v.X, Y: s.Y + v.Y}
}

func (s Vector2) Scale(scale float64) Vector2 {
	return Vector2{X: s.X * scale, Y: s.Y * scale}
}

func (s Vector2) Delta(to Vector2) Vector2 {
	return Vector2{
		X: to.X - s.X,
		Y: to.Y - s.Y,
	}
}

func (s Vector2) Length() float64 {
	return math.Sqrt(s.X*s.X + s.Y*s.Y)
}

func (s Vector2) Normalize() Vector2 {
	length := s.Length()
	if length == 0 {
		return s
	}

	return Vector2{X: s.X / length, Y: s.Y / length}
}
