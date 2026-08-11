package config

type windowCfg struct {
	Width  int
	Height int
}

func Window() windowCfg {
	return windowCfg{Width: 1920, Height: 1080}
}

func AppName() string {
	return "Asteroid"
}
