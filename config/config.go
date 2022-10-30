package config

import (
	"embed"

	"gopkg.in/yaml.v3"

	"github.com/dbut2/shortener-web/internal/web"
)

var (
	//go:embed *.yaml
	envs embed.FS
)

type Config struct {
	Web web.Config `yaml:"web"`
}

func LoadConfig(env string) (*Config, error) {
	bytes, err := envs.ReadFile(env + ".yaml")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
