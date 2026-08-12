package sound

import (
	"asteroid/assets"
	"bytes"
	"io"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const rate = 48000
var context = sync.OnceValue(func() *audio.Context { return audio.NewContext(rate) })

type Effect struct {
	bytes func() []byte
}

func NewEffect(ogg func() []byte) *Effect {
	b := sync.OnceValue(func() []byte {
		stream, err := vorbis.DecodeWithSampleRate(rate, bytes.NewReader(ogg()))
		if err != nil { 
			panic(err) 
		}

		pcm, err := io.ReadAll(stream)
		if err != nil { 
			panic(err) 
		}

    	return pcm
  	})

	return &Effect{ bytes: b }
}

func (e *Effect) Play() {
	player := context().NewPlayerFromBytes(e.bytes())
  	player.Play()
}

  var (
      Hover = NewEffect(assets.SfxHover)
      Click = NewEffect(assets.SfxClick)
  )