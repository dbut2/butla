package main

import (
	"os"

	"github.com/dbut2/butla/config"
	"github.com/dbut2/butla/internal/web"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "default"
	}

	c, err := config.LoadConfig(env)
	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	if port != "" {
		c.Web.Address = ":" + port
	}

	server, err := web.New(c.Web)
	if err != nil {
		panic(err.Error())
	}

	err = server.Run()
	if err != nil {
		panic(err.Error())
	}
}
