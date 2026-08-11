package scene

import (
	"asteroid/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(in input.Mouse) error
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

func NewSceneManager(in input.Mouse) *SceneManager {
	return &SceneManager{
		stack: []Scene{},
	}
}

func (sm *SceneManager) Push(scene Scene) {
	if len(sm.stack) > 0 {
		sm.Current().OnPause()
	}

	sm.stack = append(sm.stack, scene)
	scene.OnEnter()
}

func (sm *SceneManager) Pop() {
	if len(sm.stack) <= 0 {
		return
	}

	sm.Current().OnExit()
	sm.stack = sm.stack[:len(sm.stack)-1]

	if len(sm.stack) > 0 {
		sm.Current().OnResume()
	}
}

func (sm *SceneManager) Replace(scene Scene) {
	if len(sm.stack) > 0 {
		sm.stack[len(sm.stack)-1].OnExit()
		sm.stack[len(sm.stack)-1] = scene
	} else {
		sm.stack = append(sm.stack, scene)
	}

	sm.Current().OnEnter()
}

func (sm *SceneManager) Update(in input.Mouse) error {
	if len(sm.stack) <= 0 {
		return nil
	}

	return sm.stack[len(sm.stack)-1].Update(in)
}

func (sm *SceneManager) Draw(scene *ebiten.Image) {
	if len(sm.stack) <= 0 {
		return
	}

	sm.stack[len(sm.stack)-1].Draw(scene)
}

func (sm *SceneManager) Current() Scene {
	if len(sm.stack) <= 0 {
		return nil
	}

	return sm.stack[len(sm.stack)-1]
}
