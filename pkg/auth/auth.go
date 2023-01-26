package auth

import (
	"context"
	"github.com/dbut2/butla/pkg/models"
	"github.com/dbut2/butla/pkg/store"
	"log"
)

type Auth interface {
	Register(ctx context.Context, user models.User) error
	Login(ctx context.Context, username string, password string) (models.User, error)
}

type auth struct {
	store store.UserStore
}

var _ Auth = new(auth)

func New(store store.UserStore) Auth {
	return &auth{store: store}
}

func (a auth) Register(ctx context.Context, user models.User) error {
	if user.Username == "" {
		return ErrInvalidUser
	}
	if user.Email == "" {
		return ErrInvalidUser
	}
	if user.Name == "" {
		return ErrInvalidUser
	}
	if user.Password == "" {
		return ErrInvalidUser
	}

	_, has, err := a.store.GetUser(ctx, user.Username)
	if err != nil {
		log.Print(err.Error())
		return ErrStore
	}
	if has {
		return ErrUserExists
	}

	err = a.store.SetUser(ctx, user)
	if err != nil {
		log.Print(err.Error())
		return ErrStore
	}
	return nil
}

func (a auth) Login(ctx context.Context, username string, password string) (models.User, error) {
	if username == "" || password == "" {
		return models.User{}, ErrIncorrectAuth
	}

	user, has, err := a.store.GetUser(ctx, username)
	if err != nil {
		log.Print(err.Error())
		return models.User{}, ErrStore
	}
	if !has {
		return models.User{}, ErrIncorrectAuth
	}

	if password != user.Password {
		return models.User{}, ErrIncorrectAuth
	}

	return user, nil
}
