package config

import (
	"embed"
	"github.com/dbut2/shortener-web/internal/web"
	"github.com/dbut2/shortener-web/pkg/configs"
)

var (
	//go:embed *.yaml
	envs embed.FS
)

type Config struct {
	Web web.Config `yaml:"web"`
}

func LoadConfig(env string) (*Config, error) {
	return configs.LoadConfig[Config](envs, env)
}
