package store

import (
	"context"

	"github.com/dbut2/shortener/pkg/models"
)

type Store interface {
	Set(ctx context.Context, link models.Link) error
	Get(ctx context.Context, code string) (models.Link, error)
	Has(ctx context.Context, code string) (bool, error)
}

type CacheStore struct {
	Primary Store
	Cache   Store
}

func (c CacheStore) Set(ctx context.Context, link models.Link) error {
	err := c.Primary.Set(ctx, link)
	if err != nil {
		return err
	}
	c.Cache.Set(ctx, link)
	return nil
}

func (c CacheStore) Get(ctx context.Context, code string) (models.Link, error) {
	link, _ := c.Cache.Get(ctx, code)
	if link.Url != "" {
		return link, nil
	}
	link, err := c.Primary.Get(ctx, code)
	if err != nil {
		return models.Link{}, err
	}
	c.Cache.Set(ctx, link)
	return link, nil
}

func (c CacheStore) Has(ctx context.Context, code string) (bool, error) {
	return c.Primary.Has(ctx, code)
}
