package main

import (
	"os"

	"github.com/dbut2/butla/pkg/configs"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "default"
	}

	config, err := configs.LoadConfig[Config]()
	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.Address = ":" + port
	}

	server, err := New(config)
	if err != nil {
		panic(err.Error())
	}

	err = server.Run()
	if err != nil {
		panic(err.Error())
	}
}
