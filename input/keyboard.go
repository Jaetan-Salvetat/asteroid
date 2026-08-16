package input

import "github.com/hajimehoshi/ebiten/v2"

type Keyboard struct {
	Keys map[ebiten.Key]Phase
}

func (s *Keyboard) Next(keys []ebiten.Key) {
	keysSet := map[ebiten.Key]bool{}

	for _, key := range keys {
		keysSet[key] = true
	}

	for key, value := range keysSet {
		s.Keys[key] = s.Keys[key].Next(value)
	}

	for key := range s.Keys {
		if !keysSet[key] {
			s.Keys[key] = s.Keys[key].Next(false)
		}
	}
}

func (s Keyboard) Pressed(key ebiten.Key) bool {
	return s.Keys[key] == Pressed || s.Keys[key] == JustPressed
}
