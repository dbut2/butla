package stores

import (
	"github.com/dbut2/butla/pkg/store"
	"github.com/dbut2/butla/pkg/stores/cache"
	"github.com/dbut2/butla/pkg/stores/database"
	"github.com/dbut2/butla/pkg/stores/datastore"
	"github.com/dbut2/butla/pkg/stores/inmem"
	"github.com/dbut2/butla/pkg/stores/redis"

	"github.com/dbut2/butla/pkg/configs"
)

type Config struct {
	Database  *configs.Loader[*database.Config]  `yaml:"database"`
	Datastore *configs.Loader[*datastore.Config] `yaml:"datastore"`
	Cache     *CacheConfig                       `yaml:"cache"`
}

type CacheConfig struct {
	Redis *configs.Loader[*redis.Config] `yaml:"redis"`
}

func New(c *Config) (store.Store, error) {
	if c == nil {
		return inmem.InMem(), nil
	}

	var s store.Store

	if c.Database != nil {
		db, err := database.New(c.Database.Config)
		if err != nil {
			return nil, err
		}

		s = db
	}

	if c.Datastore != nil {
		ds, err := datastore.New(c.Datastore.Config)
		if err != nil {
			return nil, err
		}

		s = ds
	}

	if c.Cache == nil {
		return s, nil
	}

	cs := cache.Store{
		Primary: s,
		Cache:   nil,
	}

	if c.Cache.Redis != nil {
		r, err := redis.New(c.Cache.Redis.Config)
		if err != nil {
			return nil, err
		}

		cs.Cache = r
	}

	return cs, nil
}
