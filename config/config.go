package config

import (
	"flag"
	"os"

	"github.com/dbut2/shortener-web/internal/web"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Web web.Config `yaml:"web"`
}

func LoadConfig() (Config, error) {
	configPath := flag.String("config-path", "config/local.yaml", "")
	flag.Parse()

	bytes, err := os.ReadFile(*configPath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
