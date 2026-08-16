package mainmenu

import (
	"asteroid/assets"
	"asteroid/core/geo"
	"asteroid/core/render"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	baseScale           float64 = 3
	haloRestScale               = 2.16
	haloBreathAmplitude         = 0.025
	haloBreathPeriod            = 7.0
	shipTiltAmplitude           = 5 * math.Pi / 180
	shipTiltPeriod              = 17
	shipFloatAmplitude          = 20
	shipFloatPeriod             = 10
	baseShipRotation            = -60 * math.Pi / 180
)

type MenuShip struct {
	position     geo.Vector2
	anchor       geo.Vector2
	ship         *ebiten.Image
	halo         *ebiten.Image
	haloScale    float64
	shipRotation float64
	time         float64
}

func NewMenuShip(position geo.Vector2) *MenuShip {
	return &MenuShip{
		ship:      assets.Ship.Cyan(),
		halo:      assets.Shield.Cyan(),
		anchor:    position,
		position:  position,
		haloScale: baseScale,
		time:      0,
	}
}

func (s *MenuShip) Update() error {
	s.time += 1 / float64(ebiten.TPS())

	s.calculateHaloScale()
	s.calculateShipRotation()
	s.calculateLevitation()

	return nil
}

func (s *MenuShip) Draw(scene *ebiten.Image) {
	s.drawHalo(scene)
	s.drawShip(scene)
}

func (s *MenuShip) drawShip(scene *ebiten.Image) {
	opt := render.ImageOptionsCentered(s.ship.Bounds())

	opt.GeoM.Scale(baseScale, baseScale)
	opt.GeoM.Rotate(s.shipRotation)
	opt.GeoM.Translate(s.position.X, s.position.Y)
	scene.DrawImage(s.ship, opt)
}

func (s *MenuShip) drawHalo(scene *ebiten.Image) {
	opt := render.ImageOptionsCentered(s.halo.Bounds())

	opt.GeoM.Scale(s.haloScale, s.haloScale)
	opt.ColorScale.ScaleAlpha(0.16)
	opt.GeoM.Translate(s.position.X, s.position.Y)
	scene.DrawImage(s.halo, opt)
}

func (s *MenuShip) calculateHaloScale() {
	phase := 2 * math.Pi * s.time / haloBreathPeriod
	s.haloScale = haloRestScale * (1 + math.Sin(phase)*haloBreathAmplitude)
}

func (s *MenuShip) calculateShipRotation() {
	phase := 2 * math.Pi * s.time / shipTiltPeriod
	s.shipRotation = (math.Sin(phase) * shipTiltAmplitude) + baseShipRotation
}

func (s *MenuShip) calculateLevitation() {
	phase := 2 * math.Pi * s.time / shipFloatPeriod
	s.position = geo.Vector2{
		X: s.anchor.X,
		Y: s.anchor.Y + math.Sin(phase)*shipFloatAmplitude,
	}
}
