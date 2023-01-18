package main

import (
	"os"

	"github.com/dbut2/butla/configs"
	"github.com/dbut2/butla/internal/lengthener"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "default"
	}

	config, err := configs.LoadConfig(env)
	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.Web.Address = ":" + port
	}

	server, err := lengthener.New(config.Lengthener)
	if err != nil {
		panic(err.Error())
	}

	err = server.Run()
	if err != nil {
		panic(err.Error())
	}
}
