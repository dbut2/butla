package admin

import (
	"github.com/dbut2/shortener-web/pkg/configs"
	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/datastore"
)

type Config struct {
	Address string      `yaml:"address"`
	Store   storeConfig `yaml:"store"`
}

type storeConfig struct {
	Database  configs.Loader[*database.Config]  `yaml:"database"`
	Datastore configs.Loader[*datastore.Config] `yaml:"datastore"`
}
