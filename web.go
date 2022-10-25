package main

import (
	"log"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/dbut2/shortener/config"
	"github.com/dbut2/shortener/internal/web"
	"go.opencensus.io/trace"
)

func main() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)

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
