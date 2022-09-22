package redis

import (
	"context"
	"time"

	"github.com/dbut2/shortener/pkg/models"
	"github.com/dbut2/shortener/pkg/secrets"
	"github.com/dbut2/shortener/pkg/store"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	secrets.GsmResourceID `yaml:"gsmResourceID"`
	Host                  string `yaml:"host"`
	Password              string `yaml:"password"`
}

type Redis struct {
	client *redis.Client
}

var _ = store.Store(nil).(Redis)

func NewRedis(config Config) (*Redis, error) {
	err := secrets.LoadSecret(&config)
	if err != nil {
		return nil, err
	}

	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
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
