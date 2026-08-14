package config

type windowCfg struct {
	Width  float64
	Height float64
}

var window = windowCfg{Width: 1920, Height: 1080}

const AppName = "Asteroid"

func Window() windowCfg {
	return window
}
