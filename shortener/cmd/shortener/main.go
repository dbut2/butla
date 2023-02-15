package main

import (
	"os"

	"github.com/dbut2/butla/shortener/internal/server"
	"github.com/dbut2/butla/shortener/pkg/configs"
)

func main() {
	config, err := configs.LoadConfig[server.Config]()
	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.Address = ":" + port
	}

	s, err := server.New(config)
	if err != nil {
		panic(err.Error())
	}

	err = s.Run()
	if err != nil {
		panic(err.Error())
	}
}
