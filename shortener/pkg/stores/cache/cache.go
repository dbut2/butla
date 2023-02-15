package cache

import (
	"context"

	"github.com/dbut2/butla/shortener/pkg/models"
	"github.com/dbut2/butla/shortener/pkg/store"
)

type Store struct {
	Primary store.LinkStore
	Cache   store.LinkStore
}

var _ store.LinkStore = new(Store)

func (c Store) SetLink(ctx context.Context, link models.Link) error {
	err := c.Primary.SetLink(ctx, link)
	if err != nil {
		return err
	}
	_ = c.Cache.SetLink(ctx, link)
	return nil
}

func (c Store) GetLink(ctx context.Context, code string) (models.Link, bool, error) {
	link, has, _ := c.Cache.GetLink(ctx, code)
	if has {
		return link, true, nil
	}
	link, has, err := c.Primary.GetLink(ctx, code)
	if err != nil {
		return models.Link{}, false, err
	}
	if has {
		_ = c.Cache.SetLink(ctx, link)
	}
	return link, has, nil
}
