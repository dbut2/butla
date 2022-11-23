package configs

import (
	"embed"

	"gopkg.in/yaml.v3"
)

func LoadConfig[T any](envs embed.FS, env string) (*T, error) {
	bytes, err := envs.ReadFile(env + ".yaml")
	if err != nil {
		return nil, err
	}

	config := new(T)
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
