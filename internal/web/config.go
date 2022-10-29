package web

import (
	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/datastore"
	"github.com/dbut2/shortener-web/pkg/redis"
)

type Config struct {
	Address   string           `yaml:"address"`
	ShortHost string           `yaml:"shortHost"`
	Database  database.Config  `yaml:"database"`
	Datastore datastore.Config `yaml:"datastore"`
	Redis     redis.Config     `yaml:"redis"`
}
