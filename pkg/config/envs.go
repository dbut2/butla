package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Env string

var _ loader = new(Env)

func (e Env) load(c any) error {
	bytes := []byte(os.Getenv(string(e)))
	return yaml.Unmarshal(bytes, c)
}
