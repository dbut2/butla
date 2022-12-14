package web

import (
	"github.com/dbut2/butla/pkg/configs"
	"github.com/dbut2/butla/pkg/database"
	"github.com/dbut2/butla/pkg/datastore"
	"github.com/dbut2/butla/pkg/redis"
)

type Config struct {
	Address string `yaml:"address"`
	Host    host   `yaml:"host"`
	Store   Store  `yaml:"store"`
	Cache   Cache  `yaml:"cache"`
}

type host struct {
	Scheme   string `yaml:"scheme"`
	Hostname string `yaml:"hostname"`
}

type Store struct {
	Database  configs.Loader[*database.Config]  `yaml:"database"`
	Datastore configs.Loader[*datastore.Config] `yaml:"datastore"`
}

type Cache struct {
	Redis configs.Loader[*redis.Config] `yaml:"redis"`
}
