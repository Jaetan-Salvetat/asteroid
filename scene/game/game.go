package game

import (
	"asteroid/assets"
	"asteroid/input"
	"asteroid/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	navigator    scene.Navigator
	background   *ebiten.Image
	world        World
	sessionState SessionState
}

func NewScene(navigator scene.Navigator) *GameScene {
	return &GameScene{
		navigator:    navigator,
		background:   assets.Background(),
		world:        NewWorld(),
		sessionState: NewSessionState(),
	}
}

func (s *GameScene) Update(in input.Inputs) error {
	intent := NewIntentFrom(in)
	frame := s.world.Update(intent)
	s.sessionState.Update(frame)

	if s.world.IsCleared() {
		s.world.SpawnWave(s.world.ship.Circle)
	}

	return nil
}

func (s *GameScene) Draw(scene *ebiten.Image) {
	scene.DrawImage(s.background, &ebiten.DrawImageOptions{})
	s.world.Draw(scene)
}

func (s *GameScene) OnEnter() {}

func (s *GameScene) OnExit() {}

func (s *GameScene) OnPause() {}

func (s *GameScene) OnResume() {}
