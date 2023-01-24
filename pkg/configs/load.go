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

	fmt.Println(bytes)
	fmt.Println(string(bytes))

	config := new(T)
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func findBytes() ([]byte, error) {
	if str := os.Getenv("CONFIG"); str != "" {
		return []byte(str), nil
	}

	file := "config.yaml"
	if f := os.Getenv("CONFIG_FILE"); f != "" {
		file = f
	}

	_, err := os.Open(file)

	if err != os.ErrNotExist {
		if err != nil {
			return nil, err
		}

		return os.ReadFile(file)
	}

	return nil, errors.New("config not found")
}
