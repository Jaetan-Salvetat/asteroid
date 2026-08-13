package input

import (
	"asteroid/core/geo"
)

type Widget interface {
	Bounds() geo.Rect
	Place(gvector geo.Vector2)
}
