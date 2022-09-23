package config

import (
	"embed"

	"github.com/dbut2/shortener/internal/web"
	"gopkg.in/yaml.v3"
)

//go:embed *.yaml
var envs embed.FS

type Config struct {
	Web web.Config `yaml:"web"`
}

func LoadConfig(env string) (*Config, error) {
	bytes, err := envs.ReadFile(env + ".yaml")
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
