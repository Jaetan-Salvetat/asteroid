package input

import "asteroid/core/geo"

type Mouse struct {
	Cursor geo.Vector2
	State  Phase
}

func (s *Mouse) IsDown() bool {
	return s.State == Pressed || s.State == JustPressed
}

func (s *Mouse) Next(cursor geo.Vector2, isDown bool) {
	s.Cursor = cursor
	s.State = s.State.Next(isDown)
}
