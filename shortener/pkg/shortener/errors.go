package shortener

import (
	"errors"
)

var (
	ErrUnspecified   = errors.New("unspecified")
	ErrStore         = errors.New("store error")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrExpired       = errors.New("expired")
	ErrInvalidIP     = errors.New("invalid ip")
)
