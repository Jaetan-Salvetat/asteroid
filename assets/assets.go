package assets

import (
	"bytes"
	"embed"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed all:images all:fonts all:sfx
var fs embed.FS

var mainFont = sync.OnceValue(func() *text.GoTextFaceSource {
	b, err := fs.ReadFile("fonts/KenneyFuture.ttf")
	if err != nil {
		panic(err)
	}
	src, err := text.NewGoTextFaceSource(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return src
})

func Font(size float64) *text.GoTextFace {
	return &text.GoTextFace{
		Source: mainFont(),
		Size:   size,
	}
}

func lazyImage(path string) func() *ebiten.Image {
	return sync.OnceValue(func() *ebiten.Image {
		image, _, err := ebitenutil.NewImageFromFileSystem(fs, path)

		if err != nil {
			panic(err)
		}

		return image
	})
}

var (
	MenuBackground = lazyImage("images/game/background/neon_grid_1.png")
	ButtonIdle     = lazyImage("images/ui/button/normal.webp")
	ButtonActive   = lazyImage("images/ui/button/active.webp")
	ButtonHovered  = lazyImage("images/ui/button/hover.webp")
	ButtonDisabled = lazyImage("images/ui/button/disabled.webp")
)
