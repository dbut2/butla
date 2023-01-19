package datastore

import (
	"context"
	"sync"

	"cloud.google.com/go/datastore"

	"github.com/dbut2/butla/pkg/models"
	"github.com/dbut2/butla/pkg/store"
)

type Config struct {
	Project string `yaml:"project"`
}

type Datastore struct {
	client *datastore.Client
	wg     sync.WaitGroup
}

var _ store.Store = new(Datastore)

func New(c *Config) (*Datastore, error) {
	d := &Datastore{}

	d.wg.Add(1)
	go func() {
		client, err := datastore.NewClient(context.Background(), c.Project)
		if err != nil {
			panic(err.Error())
		}
		d.client = client
		d.wg.Done()
	}()

	return d, nil
}

func (d *Datastore) Set(ctx context.Context, link models.Link) error {
	d.wg.Wait()
	_, err := d.client.Put(ctx, datastore.NameKey("link", link.Code, nil), &link)
	return err
}

func (d *Datastore) Get(ctx context.Context, code string) (models.Link, bool, error) {
	d.wg.Wait()
	link := models.Link{}
	err := d.client.Get(ctx, datastore.NameKey("link", code, nil), &link)
	if err == datastore.ErrNoSuchEntity {
		return link, false, nil
	}
	if err != nil {
		return link, false, err
	}
	return link, true, nil
}
