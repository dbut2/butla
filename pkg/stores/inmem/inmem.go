package inmem

import (
	"context"
	"sync"

	"github.com/dbut2/butla/pkg/models"
	"github.com/dbut2/butla/pkg/store"
)

type inMem struct {
	links map[string]models.Link
	sync.Mutex
}

var _ store.Store = new(inMem)

func InMem() store.Store {
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
