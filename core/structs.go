package core

type Vector2 struct{ X, Y float64 }
type Rect struct{ X, Y, Width, Height float64 }

func (r *Rect) Contains(p Vector2) bool {
	xEnd := r.X + r.Width
	yEnd := r.Y + r.Height

	return  p.X >= r.X && p.X <= xEnd && p.Y >= r.Y && p.Y <= yEnd
}
