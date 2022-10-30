package web

import (
	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/datastore"
	"github.com/dbut2/shortener-web/pkg/redis"
)

type Config struct {
	Address   string    `yaml:"address"`
	ShortHost ShortHost `yaml:"shortHost"`
	Store     Store     `yaml:"store"`
	Cache     Cache     `yaml:"cache"`
}

type ShortHost struct {
	Scheme string `yaml:"scheme" json:"scheme" `
	URL    string `yaml:"host"`
}

type Store struct {
	Database  *database.Config  `yaml:"database"`
	Datastore *datastore.Config `yaml:"datastore"`
}

type Cache struct {
	Redis *redis.Config `yaml:"redis"`
}
