package store

import (
	"context"
	"log"

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
	has, _ := c.Cache.Has(ctx, code)
	if has {
		return true, nil
	}
	return c.Primary.Has(ctx, code)
}

func Log(store Store, name string) Store {
	return logStore{
		name:  name,
		store: store,
	}
}

type logStore struct {
	name  string
	store Store
}

func (l logStore) Set(ctx context.Context, link models.Link) error {
	log.Printf("%s: set", l.name)
	defer log.Printf("%s: setted", l.name)
	return l.store.Set(ctx, link)
}

func (l logStore) Get(ctx context.Context, code string) (models.Link, error) {
	log.Printf("%s: get", l.name)
	defer log.Printf("%s: getted", l.name)
	return l.store.Get(ctx, code)
}

func (l logStore) Has(ctx context.Context, code string) (bool, error) {
	log.Printf("%s: has", l.name)
	defer log.Printf("%s: hassed", l.name)
	return l.store.Has(ctx, code)
}
