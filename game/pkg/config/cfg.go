package config

type Config struct {
	Width  int
	Height int
}

func (c Config) New() *Config {

	c.Width = 100
	c.Height = 50

	return &c
}
