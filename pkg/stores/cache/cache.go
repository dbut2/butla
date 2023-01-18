package cache

import (
	"context"

	"github.com/dbut2/butla/pkg/models"
	"github.com/dbut2/butla/pkg/store"
)

type Store struct {
	Primary store.Store
	Cache   store.Store
}

var _ store.Store = new(Store)

func (c Store) Set(ctx context.Context, link models.Link) error {
	err := c.Primary.Set(ctx, link)
	if err != nil {
		return err
	}
	_ = c.Cache.Set(ctx, link)
	return nil
}

func (c Store) Get(ctx context.Context, code string) (models.Link, bool, error) {
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
