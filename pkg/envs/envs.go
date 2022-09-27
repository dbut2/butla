package envs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Env string

func (e Env) getEnv() string {
	return os.Getenv(string(e))
}

type enver interface {
	getEnv() string
}

func LoadEnv(config enver) error {
	bytes := []byte(config.getEnv())
	return yaml.Unmarshal(bytes, config)
}
