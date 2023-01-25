package configs

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig[T any]() (*T, error) {
	bytes, err := findBytes()
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

func findBytes() ([]byte, error) {
	file := "config.yaml"

	if e := os.Getenv("ENV"); e != "" {
		f := fmt.Sprintf("config/%s.yaml", e)
		if _, err := os.Stat(f); err == nil {
			file = f
		}
	}

	if f := os.Getenv("CONFIG_FILE"); f != "" {
		file = f
	}

	if _, err := os.Stat(file); err != nil {
		return nil, errors.New("config file not found")
	}

	return os.ReadFile(file)
}
