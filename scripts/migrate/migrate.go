package main

import (
	"context"
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/dbut2/butla/pkg/configs"
	"github.com/dbut2/butla/pkg/stores/database"
	"github.com/dbut2/butla/pkg/stores/datastore"
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
		log.Fatal(err.Error())
	}

	db, err := database.New(c.Database.Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	ds, err := datastore.New(c.Datastore.Config)
	if err != nil {
		log.Fatal(err.Error())
	}

	links, err := db.GetAll(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, link := range links {
		err = ds.SetLink(context.Background(), link)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
