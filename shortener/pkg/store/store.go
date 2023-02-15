package store

import (
	"context"

	"github.com/dbut2/butla/shortener/pkg/models"
)

type Store interface {
	LinkStore
	UserStore
}

type LinkStore interface {
	SetLink(ctx context.Context, link models.Link) error
	GetLink(ctx context.Context, code string) (models.Link, bool, error)
}

type UserStore interface {
	SetUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, username string) (models.User, bool, error)
}

type GenericStore[T any] interface {
	Set(ctx context.Context, val T) error
	Get(ctx context.Context, key string) (T, bool, error)
}
