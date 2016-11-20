package config

const configFile string = ".tasks.toml"

type Config struct {
	Path      string
	Formatter string
}

func New() (*Config, error) {
	return &Config{Path: "defaultPath", Formatter: "ascii"}, nil
}
