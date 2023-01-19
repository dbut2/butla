package configs

import (
	"embed"

	"github.com/dbut2/butla/internal/lengthener"
	"github.com/dbut2/butla/internal/web"
	"github.com/dbut2/butla/pkg/configs"
)

var (
	//go:embed *.yaml
	envs embed.FS
)

type Config struct {
	Lengthener *lengthener.Config `yaml:"lengthener"`
	Web        *web.Config        `yaml:"web"`
}

func LoadConfig(env string) (*Config, error) {
	return configs.LoadConfig[Config](envs, env)
}
