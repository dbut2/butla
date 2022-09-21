package web

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/dbut2/shortener/pkg/database"
	"github.com/dbut2/shortener/pkg/models"
	"github.com/dbut2/shortener/pkg/redis"
	"github.com/gin-gonic/gin"

	"github.com/dbut2/shortener/pkg/store"
)

type Server struct {
	address   string
	shortHost string
	store     store.Store
}

type CacheLog struct {
	cache store.Cache
}

func (c CacheLog) Set(ctx context.Context, link models.Link) {
	log.Print("cache: set")
	c.cache.Set(ctx, link)
}

func (c CacheLog) Get(ctx context.Context, code string) models.Link {
	log.Print("cache: get")
	return c.cache.Get(ctx, code)
}

func (c CacheLog) Has(ctx context.Context, code string) bool {
	log.Print("cache: has")
	return c.cache.Has(ctx, code)
}

type StoreLog struct {
	store store.Store
}

func (s StoreLog) Set(ctx context.Context, link models.Link) error {
	log.Print("store: set")
	return s.store.Set(ctx, link)
}

func (s StoreLog) Get(ctx context.Context, code string) (models.Link, error) {
	log.Print("store: get")
	return s.store.Get(ctx, code)
}

func (s StoreLog) Has(ctx context.Context, code string) (bool, error) {
	log.Print("store: has")
	return s.store.Has(ctx, code)
}

func New(config Config) (*Server, error) {
	db, err := database.NewDatabase(config.Database)
	if err != nil {
		return nil, err
	}

	r, err := redis.NewRedis(config.Redis)
	if err != nil {
		return nil, err
	}

	s := store.CacheStore{
		Primary: StoreLog{db},
		Cache:   CacheLog{r},
	}

	return &Server{
		address:   config.Address,
		shortHost: config.ShortHost,
		store:     s,
	}, nil
}

func (s *Server) Run() error {
	r := gin.Default()

	r.GET("/shorten", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", index)
	})

	r.POST("/shorten", func(c *gin.Context) {
		b := struct {
			Url string `json:"url"`
		}{}

		var code string
		for {
			code = randomCode(6)
			has, err := s.store.Has(c, code)
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			if !has {
				break
			}
		}

		link := models.Link{
			Code: code,
			Url:  b.Url,
			Expiry: models.NullTime{
				Valid: true,
				Value: time.Now().Add(time.Minute * 10),
			},
			IP: models.NullString{
				Valid: true,
				Value: c.ClientIP(),
			},
		}

		err := s.store.Set(c, link)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, struct {
			Link string `json:"link"`
		}{
			Link: fmt.Sprintf("%s/%s", s.shortHost, code),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "default")
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")

		link, err := s.store.Get(c, code)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if link.Expiry.Valid && link.Expiry.Value.Before(time.Now()) {
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("link expired"))
			return
		}

		if link.IP.Valid && link.IP.Value != c.ClientIP() {
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("incorrect ip"))
			return
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	})

	return r.Run(s.address)
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
