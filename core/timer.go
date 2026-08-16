package core

import "github.com/hajimehoshi/ebiten/v2"

type Timer struct {
	expiry    float64
	baseValue float64
}

func NewTimer(expiry float64) Timer {
	return Timer{expiry: expiry, baseValue: expiry}
}

func (s *Timer) Update() {
	s.expiry -= 1 / float64(ebiten.TPS())
}

func (s *Timer) IsExpired() bool {
	return s.expiry <= 0
}

func (s *Timer) Reset() {
	s.expiry = s.baseValue
}
