package input

import "asteroid/core/geo"

type Phase uint8

const (
	Unknown Phase = iota
	Pressed
	JustPressed
	Released
	JustReleased
)

type Mouse struct {
	Cursor geo.Vector2
	State  Phase
}

func (m *Mouse) IsDown() bool {
	return m.State == Pressed || m.State == JustPressed
}

func (m *Mouse) Next(cursor geo.Vector2, isDown bool) {
	m.Cursor = cursor

	if m.State == Unknown {
		m.State = Released
		return
	}

	if isDown {
		if m.IsDown() {
			m.State = Pressed
			return
		}

		m.State = JustPressed
		return
	}

	if m.IsDown() {
		m.State = JustReleased
		return
	}

	m.State = Released
}
