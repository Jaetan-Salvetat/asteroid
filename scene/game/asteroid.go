package game

import (
	"asteroid/assets"
	"asteroid/config"
	"asteroid/core/geo"
	"asteroid/core/render"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const maxAsteroid = 8

type AsteroidSize int

const (
	AsteroidSizeSmall AsteroidSize = iota
	AsteroidSizeMedium
	AsteroidSizeLarge
)

type Asteroid struct {
	asset  *ebiten.Image
	size   AsteroidSize
	circle geo.Circle
}

func NewAsteroid(position geo.Vector2) Asteroid {
	asset := assets.Asteroid.Large()
	return Asteroid{
		asset: asset,
		size:  AsteroidSizeLarge,
		circle: geo.Circle{
			Center: position,
			Radius: float64(asset.Bounds().Dx()) / 2,
		},
	}
}

func (s *Asteroid) Draw(scene *ebiten.Image) {
	opt := render.ImageOptionsCentered(s.asset.Bounds())
	opt.GeoM.Translate(s.circle.Center.X, s.circle.Center.Y)

	scene.DrawImage(s.asset, opt)
}

type Asteroids []Asteroid

func (s *Asteroids) Draw(scene *ebiten.Image) {
	for _, asteroid := range *s {
		asteroid.Draw(scene)
	}
}

func (s *Asteroids) SpawnWaved(ship geo.Circle) {
	window := config.Window()
	margin := float64(assets.Asteroid.Large().Bounds().Dx()) / 2
	asteroidsSpawned := 0

	for asteroidsSpawned < maxAsteroid {
		position := geo.Vector2{
			X: float64(rand.Intn(int(window.Width-margin*2))) + margin,
			Y: float64(rand.Intn(int(window.Height-margin*2))) + margin,
		}
		asteroid := NewAsteroid(position)

		if s.overlaps(asteroid.circle, ship) {
			continue
		}

		*s = append(*s, asteroid)
		asteroidsSpawned++
	}
}

func (s *Asteroids) overlaps(other geo.Circle, ship geo.Circle) bool {
	if ship.Intersects(other) {
		return true
	}

	for _, asteroid := range *s {
		if asteroid.circle.Intersects(other) {
			return true
		}
	}

	return false
}
