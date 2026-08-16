package assets

import (
	"bytes"
	"io"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const rate = 48000

var ctx = sync.OnceValue(func() *audio.Context { return audio.NewContext(rate) })

type Effect struct {
	bytes func() []byte
}

func newEffect(path string) Effect {
	b := sync.OnceValue(func() []byte {
		soundBytes, err := fs.ReadFile(path)
		if err != nil {
			panic(err)
		}

		stream, err := vorbis.DecodeWithSampleRate(rate, bytes.NewReader(soundBytes))
		if err != nil {
			panic(err)
		}

		pcm, err := io.ReadAll(stream)
		if err != nil {
			panic(err)
		}

		return pcm
	})

	return Effect{bytes: b}
}

func (e Effect) Play() {
	player := ctx().NewPlayerFromBytes(e.bytes())
	player.Play()
}

var (
	Hover = newEffect("sfx/hover.ogg")
	Click = newEffect("sfx/click.ogg")
)
