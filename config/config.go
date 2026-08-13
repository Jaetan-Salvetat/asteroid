package config

type windowCfg struct {
	Width  float64
	Height float64
}

func Window() windowCfg {
	return windowCfg{Width: 1920, Height: 1080}
}

func AppName() string {
	return "Asteroid"
}
