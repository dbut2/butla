package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/dbut2/butla/shortener/pkg/models"
	"github.com/dbut2/butla/shortener/pkg/store"
)

type Config struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

type Redis struct {
	client *redis.Client
}

var _ store.LinkStore = new(Redis)

func New(c *Config) (*Redis, error) {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
	})}, nil
}

func (r Redis) SetLink(ctx context.Context, link models.Link) error {
	expiry := time.Hour * 24 * 7
	if link.Expiry.Valid {
		expiry = time.Until(link.Expiry.Value)
	}
	return r.client.Set(ctx, link.Code, link, expiry).Err()
}

func (r Redis) GetLink(ctx context.Context, code string) (models.Link, bool, error) {
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
