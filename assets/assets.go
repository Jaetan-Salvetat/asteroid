package assets

import (
	"embed"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed images
var fs embed.FS

func lazyLoad(path string) func() *ebiten.Image {
	return sync.OnceValue(func() *ebiten.Image {
		image, _, err := ebitenutil.NewImageFromFileSystem(fs, path)

		if err != nil {
			panic(err)
		}

		return image
	})
}

var (
	MenuBackground = lazyLoad("images/game/background/neon_grid_1.png")
)
