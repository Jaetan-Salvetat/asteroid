package assets

import (
	"bytes"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var fontSource = sync.OnceValue(func() *text.GoTextFaceSource {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))

	if err != nil {
		panic(err)
	}
	
	return source
})

func Font(size float64) *text.GoTextFace {
	return &text.GoTextFace{Source: fontSource(), Size: size}
}