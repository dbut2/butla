package auth

import "errors"

var (
	ErrUnspecified   = errors.New("unspecified")
	ErrUserExists    = errors.New("user already exists")
	ErrUserNotExists = errors.New("user not exists")
	ErrIncorrectAuth = errors.New("incorrect auth")
	ErrInvalidUser   = errors.New("invalid user")
	ErrStore         = errors.New("store error")
)
