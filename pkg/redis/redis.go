package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/dbut2/shortener-web/pkg/config"
	"github.com/dbut2/shortener-web/pkg/models"
	"github.com/dbut2/shortener-web/pkg/store"
)

type Config struct {
	config.Loader `yaml:",inline"`
	Host          string `yaml:"host"`
	Password      string `yaml:"password"`
}

type Redis struct {
	client *redis.Client
}

var _ store.Store = new(Redis)

func NewRedis(c Config) (*Redis, error) {
	err := config.Load(&c)
	if err != nil {
		return nil, err
	}

	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
	})}, nil
}

func (r Redis) Set(ctx context.Context, link models.Link) error {
	expiry := time.Hour * 24 * 7
	if link.Expiry.Valid {
		expiry = time.Until(link.Expiry.Value)
	}
	return r.client.Set(ctx, link.Code, link, expiry).Err()
}

func (r Redis) Get(ctx context.Context, code string) (models.Link, bool, error) {
	if r.client.Exists(ctx, code).Val() == 0 {
		return models.Link{}, false, nil
	}

	g := r.client.Get(ctx, code)
	if g.Err() != nil {
		return models.Link{}, false, g.Err()
	}

	var link models.Link
	return link, true, g.Scan(&link)
}
