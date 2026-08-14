package assets

import (
	_ "image/png"

	_ "golang.org/x/image/webp"

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

func lazyBytes(path string) func() []byte {
	return sync.OnceValue(func() []byte {
		b, err := fs.ReadFile(path)
		if err != nil {
			panic(err)
		}

		return b
	})
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
	SfxHover = lazyBytes("sfx/hover.ogg")
	SfxClick = lazyBytes("sfx/click.ogg")

	MenuBackground = lazyImage("images/game/background/neon_grid_1.png")
	ShipCyan       = lazyImage("images/game/ship/01_cyan.png")
	ShieldRing1    = lazyImage("images/game/shield/ring_2.png")
	ButtonIdle     = lazyImage("images/ui/button/normal.webp")
	ButtonActive   = lazyImage("images/ui/button/active.webp")
	ButtonHovered  = lazyImage("images/ui/button/hover.webp")
	ButtonDisabled = lazyImage("images/ui/button/disabled.webp")

	AsteroidSmall01  = lazyImage("images/game/asteroid/small_01.png")
	AsteroidSmall03  = lazyImage("images/game/asteroid/small_03.png")
	AsteroidSmall05  = lazyImage("images/game/asteroid/small_05.png")
	AsteroidMedium01 = lazyImage("images/game/asteroid/medium_01.png")
	AsteroidMedium03 = lazyImage("images/game/asteroid/medium_03.png")
	AsteroidMedium05 = lazyImage("images/game/asteroid/medium_05.png")
	AsteroidLarge01  = lazyImage("images/game/asteroid/large_01.png")
	AsteroidLarge03  = lazyImage("images/game/asteroid/large_03.png")
	AsteroidLarge05  = lazyImage("images/game/asteroid/large_05.png")
)
