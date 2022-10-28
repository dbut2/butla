package shortener

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"
	"time"

	"github.com/dbut2/shortener/pkg/models"
	"github.com/dbut2/shortener/pkg/store"
	"go.opencensus.io/trace"
)

type Shortener interface {
	Shorten(ctx context.Context, url string, metadata ...Metadata) (models.Link, error)
	ShortenCode(ctx context.Context, url string, code string, metadata ...Metadata) (models.Link, error)
	Lengthen(ctx context.Context, code string, metadata ...Metadata) (models.Link, error)
}

type md struct {
	expiry models.NullTime
	ip     models.NullString
}

type Metadata func(md md) md

func WithExpiry(expiry time.Time) Metadata {
	return func(md md) md {
		md.expiry = models.NullTime{
			Valid: true,
			Value: expiry,
		}
		return md
	}
}

func WithIP(ip string) Metadata {
	return func(md md) md {
		md.ip = models.NullString{
			Valid: true,
			Value: ip,
		}
		return md
	}
}

type shortener struct {
	store store.Store
}

func New(store store.Store) Shortener {
	return shortener{store: store}
}

func (d shortener) Shorten(ctx context.Context, url string, metadata ...Metadata) (models.Link, error) {
	ctx, span := trace.StartSpan(ctx, "shorten.Shorten")
	defer span.End()
	var code string
	for {
		code = randomCode(6)
		_, has, err := d.store.Get(ctx, code)
		if err != nil {
			log.Print(err.Error())
			return models.Link{}, ErrStore
		}
		if !has {
			break
		}
	}
	return d.ShortenCode(ctx, url, code, metadata...)
}

func randomCode(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz1234567890"
	code := ""
	for len(code) < length {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err.Error())
		}

		code += string(chars[n.Int64()])
	}

	return code
}

func (d shortener) ShortenCode(ctx context.Context, url string, code string, metadata ...Metadata) (models.Link, error) {
	ctx, span := trace.StartSpan(ctx, "shorten.ShortenCode")
	defer span.End()
	md := md{}
	for _, m := range metadata {
		md = m(md)
	}

	if url == "" {
		return models.Link{}, ErrUnspecified
	}

	_, has, err := d.store.Get(ctx, code)
	if err != nil {
		return models.Link{}, ErrStore
	}
	if has {
		return models.Link{}, ErrAlreadyExists
	}

	link := models.Link{
		Code:   code,
		Url:    url,
		Expiry: md.expiry,
		IP:     md.ip,
	}

	err = d.store.Set(ctx, link)
	if err != nil {
		log.Print(err.Error())
		return models.Link{}, ErrStore
	}

	return link, nil
}

func (d shortener) Lengthen(ctx context.Context, code string, metadata ...Metadata) (models.Link, error) {
	ctx, span := trace.StartSpan(ctx, "shorten.Length")
	defer span.End()
	md := md{}
	for _, m := range metadata {
		md = m(md)
	}

	link, has, err := d.store.Get(ctx, code)
	if err != nil {
		log.Print(err.Error())
		return models.Link{}, ErrStore
	}
	if !has {
		return models.Link{}, ErrNotFound
	}

	if link.Expiry.Valid && link.Expiry.Value.Before(time.Now()) {
		return models.Link{}, ErrExpired
	}

	if link.IP.Valid && (!md.ip.Valid || link.IP.Value != md.ip.Value) {
		return models.Link{}, ErrInvalidIP
	}

	return link, nil
}
