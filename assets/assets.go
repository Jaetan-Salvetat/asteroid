package assets

import (
	"embed"
)

//go:embed all:images all:fonts all:sfx
var fs embed.FS
