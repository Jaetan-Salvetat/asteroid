package assets

import (
	_ "golang.org/x/image/webp"

	"bytes"
	"embed"
	"sync"

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
