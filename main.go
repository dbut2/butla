package main

import (
	"os"

	"github.com/dbut2/shortener-web/config"
	"github.com/dbut2/shortener-web/internal/web"
)

func main() {
	c, err := config.LoadConfig()
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
