package models

import (
	"time"
)

type NullAny[T any] struct {
	Valid bool
	Value T
}

type NullTime NullAny[time.Time]

type NullString NullAny[string]
