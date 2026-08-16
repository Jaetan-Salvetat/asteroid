package input

import "github.com/hajimehoshi/ebiten/v2"

type Phase int

const (
	Released Phase = iota
	JustReleased
	JustPressed
	Pressed
)

func (s Phase) IsDown() bool {
	return s == Pressed || s == JustPressed
}

func (s Phase) Next(isDown bool) Phase {
	if isDown {
		if s.IsDown() {
			return Pressed
		}

		return JustPressed
	}

	if s.IsDown() {
		return JustReleased
	}

	return Released
}

type Inputs struct {
	Mouse    Mouse
	Keyboard Keyboard
}

func NewInputs() Inputs {
	return Inputs{
		Keyboard: Keyboard{Keys: map[ebiten.Key]Phase{}},
	}
}
