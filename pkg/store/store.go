package store

import (
	"context"
	"log"
	"sync"

	"github.com/dbut2/shortener-web/pkg/models"
)

type Store interface {
	Set(ctx context.Context, link models.Link) error
	Get(ctx context.Context, code string) (models.Link, bool, error)
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
	_ = c.Cache.Set(ctx, link)
	return nil
}

func (c CacheStore) Get(ctx context.Context, code string) (models.Link, bool, error) {
	link, has, _ := c.Cache.Get(ctx, code)
	if has {
		return link, true, nil
	}
	link, has, err := c.Primary.Get(ctx, code)
	if err != nil {
		return models.Link{}, false, err
	}
	if has {
		_ = c.Cache.Set(ctx, link)
	}
	return link, has, nil
}

type inMem struct {
	links map[string]models.Link
	sync.Mutex
}

func InMem() Store {
	return &inMem{
		links: make(map[string]models.Link),
	}
}

func (i *inMem) Set(ctx context.Context, link models.Link) error {
	i.Lock()
	i.links[link.Code] = link
	i.Unlock()
	return nil
}

func (i *inMem) Get(ctx context.Context, code string) (models.Link, bool, error) {
	i.Lock()
	link, has := i.links[code]
	i.Unlock()
	return link, has, nil
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

func (l logStore) Get(ctx context.Context, code string) (models.Link, bool, error) {
	log.Printf("%s: get", l.name)
	defer log.Printf("%s: getted", l.name)
	return l.store.Get(ctx, code)
}
