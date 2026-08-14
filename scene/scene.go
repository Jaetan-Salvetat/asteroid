package scene

import (
	"asteroid/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(in input.Inputs) error
	Draw(*ebiten.Image)
	OnEnter()
	OnExit()
	OnPause()
	OnResume()
}

type Navigator interface {
	Push(Scene)
	Pop()
	Replace(Scene)
}

type SceneManager struct {
	stack []Scene
}

func NewManager() *SceneManager {
	return &SceneManager{
		stack: []Scene{},
	}
}

func (s *SceneManager) Push(scene Scene) {
	if len(s.stack) > 0 {
		s.Current().OnPause()
	}

	s.stack = append(s.stack, scene)
	scene.OnEnter()
}

func (s *SceneManager) Pop() {
	if len(s.stack) <= 1 {
		return
	}

	last := len(s.stack) - 1
	s.stack[last].OnExit()
	s.stack[last] = nil
	s.stack = s.stack[:len(s.stack)-1]

	if len(s.stack) > 0 {
		s.Current().OnResume()
	}
}

func (s *SceneManager) Replace(scene Scene) {
	if len(s.stack) > 0 {
		s.stack[len(s.stack)-1].OnExit()
		s.stack[len(s.stack)-1] = scene
	} else {
		s.stack = append(s.stack, scene)
	}

	s.Current().OnEnter()
}

func (s *SceneManager) Update(in input.Inputs) error {
	if len(s.stack) <= 0 {
		return nil
	}

	return s.stack[len(s.stack)-1].Update(in)
}

func (s *SceneManager) Draw(scene *ebiten.Image) {
	if len(s.stack) <= 0 {
		return
	}

	s.stack[len(s.stack)-1].Draw(scene)
}

func (s *SceneManager) Current() Scene {
	if len(s.stack) <= 0 {
		return nil
	}

	return s.stack[len(s.stack)-1]
}
