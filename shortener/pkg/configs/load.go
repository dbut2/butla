package configs

import (
	"errors"
	"fmt"
	"log"
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

	if env := os.Getenv("ENV"); env != "" {
		f := fmt.Sprintf("configs/%s.yaml", env)
		if _, err := os.Stat(f); err == nil {
			file = f
		}
	}

	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" {
		file = configFile
	}

	log.Printf("config: %s", file)

	if _, err := os.Stat(file); err != nil {
		return nil, errors.New("config file not found")
	}

	return os.ReadFile(file)
}
