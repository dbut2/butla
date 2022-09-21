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

type Cache interface {
	Set(ctx context.Context, link models.Link)
	Get(ctx context.Context, code string) models.Link
	Has(ctx context.Context, code string) bool
}

type CacheStore struct {
	Primary Store
	Cache   Cache
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
	if c.Cache.Has(ctx, code) {
		return c.Cache.Get(ctx, code), nil
	}
	link, err := c.Primary.Get(ctx, code)
	if err != nil {
		return models.Link{}, err
	}
	c.Cache.Set(ctx, link)
	return link, nil
}

func (c CacheStore) Has(ctx context.Context, code string) (bool, error) {
	if c.Cache.Has(ctx, code) {
		return true, nil
	}
	return c.Primary.Has(ctx, code)
}
