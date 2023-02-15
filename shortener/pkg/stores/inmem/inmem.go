package inmem

import (
	"context"
	"sync"

	"github.com/dbut2/butla/shortener/pkg/models"
	"github.com/dbut2/butla/shortener/pkg/store"
)

type inMem struct {
	links map[string]models.Link
	users map[string]models.User
	sync.Mutex
}

var _ store.LinkStore = new(inMem)
var _ store.UserStore = new(inMem)

func InMem() store.Store {
	return &inMem{
		links: make(map[string]models.Link),
		users: make(map[string]models.User),
	}
}

func (i *inMem) SetLink(ctx context.Context, link models.Link) error {
	i.Lock()
	i.links[link.Code] = link
	i.Unlock()
	return nil
}

func (i *inMem) GetLink(ctx context.Context, code string) (models.Link, bool, error) {
	i.Lock()
	link, has := i.links[code]
	i.Unlock()
	return link, has, nil
}

func (i *inMem) SetUser(ctx context.Context, user models.User) error {
	i.Lock()
	i.users[user.Username] = user
	i.Unlock()
	return nil
}

func (i *inMem) GetUser(ctx context.Context, username string) (models.User, bool, error) {
	i.Lock()
	user, has := i.users[username]
	i.Unlock()
	return user, has, nil
}
