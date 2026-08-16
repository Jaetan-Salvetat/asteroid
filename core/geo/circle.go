package geo

type Circle struct {
	Center Vector2
	Radius float64
}

func (s Circle) Intersects(other Circle) bool {
	delta := other.Center.Delta(s.Center)
	radius := s.Radius + other.Radius

	return delta.Length() <= radius
}
