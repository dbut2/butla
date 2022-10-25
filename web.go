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
	err := startStackdriver()

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	c, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port != "" {
		c.Web.Address = ":" + port
	}

	server, err := web.New(c.Web)
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func startStackdriver() error {
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if project == "" {
		return nil
	}
	exporter, err := stackdriver.NewExporter(stackdriver.Options{ProjectID: project})
	if err != nil {
		return err
	}
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	return nil
}
