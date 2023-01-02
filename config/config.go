package config

import (
	"embed"

	"github.com/dbut2/butla/internal/api"
	"github.com/dbut2/butla/internal/web"
	"github.com/dbut2/butla/pkg/configs"
)

var (
	//go:embed *.yaml
	envs embed.FS
)

type Config struct {
	Web web.Config `yaml:"web"`
	API api.Config `yaml:"api"`
}

func LoadConfig(env string) (*Config, error) {
	return configs.LoadConfig[Config](envs, env)
}
