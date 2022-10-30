package main

import (
	"os"

	"github.com/dbut2/shortener-web/config"
	"github.com/dbut2/shortener-web/internal/admin"
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

	server, err := admin.New(c.Admin)
	if err != nil {
		panic(err.Error())
	}

	err = server.Run()
	if err != nil {
		panic(err.Error())
	}
}
