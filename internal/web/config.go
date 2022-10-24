package web

import (
	"github.com/dbut2/shortener/pkg/database"
	"github.com/dbut2/shortener/pkg/datastore"
	"github.com/dbut2/shortener/pkg/redis"
)

type Config struct {
	Address   string           `yaml:"address"`
	ShortHost string           `yaml:"shortHost"`
	Database  database.Config  `yaml:"database"`
	Datastore datastore.Config `yaml:"datastore"`
	Redis     redis.Config     `yaml:"redis"`
}
