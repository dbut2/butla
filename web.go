package main

import (
	"os"

	"github.com/dbut2/shortener/config"
	"github.com/dbut2/shortener/internal/web"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
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
