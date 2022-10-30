package main

import (
	"context"
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/dbut2/shortener-web/pkg/configs"
	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/datastore"
)

type Config struct {
	Database  configs.Loader[*database.Config]  `yaml:"database"`
	Datastore configs.Loader[*datastore.Config] `yaml:"datastore"`
}

//go:embed migrate.yaml
var config []byte

func main() {
	var c Config
	err := yaml.Unmarshal(config, &c)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db, err := database.New(*c.Database.C)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ds, err := datastore.New(*c.Datastore.C)
	if err != nil {
		log.Fatalln(err.Error())
	}

	links, err := db.GetAll(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, link := range links {
		err = ds.Set(context.Background(), link)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
