package store

import (
	"context"

	"github.com/dbut2/butla/pkg/models"
)

type Store interface {
	Set(ctx context.Context, link models.Link) error
	Get(ctx context.Context, code string) (models.Link, bool, error)
}
