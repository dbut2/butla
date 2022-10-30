package configs

import (
	"os"
)

type Env string

func (e Env) load() ([]byte, error) {
	bytes := []byte(os.Getenv(string(e)))
	return bytes, nil
}
