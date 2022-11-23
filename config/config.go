package config

import (
	"embed"

	"github.com/but-la/pkg/configs"

	"github.com/but-la/web/internal/web"
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
