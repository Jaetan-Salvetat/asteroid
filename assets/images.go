package assets

import (
	"fmt"
	"math/rand/v2"
	"sync"

	_ "golang.org/x/image/webp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadImage(path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFileSystem(fs, path)

	if err != nil {
		panic(err)
	}

	return image
}

func lazyImage(path string) func() *ebiten.Image {
	return sync.OnceValue(func() *ebiten.Image {
		return loadImage(path)
	})
}

type button struct {
	Idle     func() *ebiten.Image
	Active   func() *ebiten.Image
	Hover    func() *ebiten.Image
	Disabled func() *ebiten.Image
}

type ship struct {
	Cyan   func() *ebiten.Image
	Green  func() *ebiten.Image
	Pink   func() *ebiten.Image
	Purple func() *ebiten.Image
	Blue   func() *ebiten.Image
	Yellow func() *ebiten.Image
}

type projectile struct {
	Bullet func() *ebiten.Image
	Laser  func() *ebiten.Image
	UFO    func() *ebiten.Image
}

type shield struct {
	Blue   func() *ebiten.Image
	Cyan   func() *ebiten.Image
	Purple func() *ebiten.Image
}

type asteroid struct {
	Small  func() []*ebiten.Image
	Medium func() []*ebiten.Image
	Large  func() []*ebiten.Image
}

const maxAsteroid = 6

type AsteroidSize string

const (
	Small  AsteroidSize = "small"
	Medium AsteroidSize = "medium"
	Large  AsteroidSize = "large"
)

func lazyLoadAsteroids(size AsteroidSize) func() []*ebiten.Image {
	return sync.OnceValue(func() []*ebiten.Image {
		images := []*ebiten.Image{}

		for i := range maxAsteroid {
			image := loadImage(fmt.Sprintf("images/game/asteroid/%s_0%d.webp", size, i+1))
			images = append(images, image)
		}

		return images
	})
}

var (
	Background = lazyImage("images/game/background/neon_grid_1.webp")

	Button = button{
		Idle:     lazyImage("images/ui/button/normal.webp"),
		Active:   lazyImage("images/ui/button/active.webp"),
		Hover:    lazyImage("images/ui/button/hover.webp"),
		Disabled: lazyImage("images/ui/button/disabled.webp"),
	}

	Ship = ship{
		Cyan:   lazyImage("images/game/ship/01_cyan.webp"),
		Green:  lazyImage("images/game/ship/02_pink.webp"),
		Pink:   lazyImage("images/game/ship/03_green.webp"),
		Purple: lazyImage("images/game/ship/04_purple.webp"),
		Blue:   lazyImage("images/game/ship/05_blue.webp"),
		Yellow: lazyImage("images/game/ship/06_yellow.webp"),
	}

	Projectile = projectile{
		Bullet: lazyImage("images/game/projectile/bullet_cyan.webp"),
		Laser:  lazyImage("images/game/projectile/laser_magenta.webp"),
		UFO:    lazyImage("images/game/projectile/ufo_shot_red.webp"),
	}

	Shield = shield{
		Blue:   lazyImage("images/game/shield/ring_1.webp"),
		Cyan:   lazyImage("images/game/shield/ring_2.webp"),
		Purple: lazyImage("images/game/shield/ring_3.webp"),
	}

	asteroids = asteroid{
		Small:  lazyLoadAsteroids(Small),
		Medium: lazyLoadAsteroids(Medium),
		Large:  lazyLoadAsteroids(Large),
	}

	// return a random asteroid image for the given size
	Asteroid = func(size AsteroidSize) *ebiten.Image {
		index := rand.IntN(maxAsteroid - 1)

		switch size {
		case Medium:
			return asteroids.Medium()[index]
		case Large:
			return asteroids.Large()[index]
		default:
			return asteroids.Small()[index]
		}
	}
)
