package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/shortener/pkg/database"
	"github.com/dbut2/shortener/pkg/redis"
	"github.com/dbut2/shortener/pkg/shortener"
	"github.com/dbut2/shortener/pkg/store"
)

type Server struct {
	address   string
	shortHost string
	shortener shortener.Shortener
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
		Primary: store.Log(db, "primary"),
		Cache:   store.Log(r, "cache"),
	}

	return &Server{
		address:   config.Address,
		shortHost: config.ShortHost,
		shortener: shortener.New(store.Log(s, "main")),
	}, nil
}

func (s *Server) Run() error {
	r := gin.Default()

	r.GET("/shorten", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", index)
	})

	r.GET("/404", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", e404)
	})

	r.POST("/shorten", func(c *gin.Context) {
		b := struct {
			Url string `json:"url"`
		}{}
		err := c.BindJSON(&b)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		link, err := s.shortener.Shorten(c, b.Url, shortener.WithExpiry(time.Now().Add(time.Minute*10)), shortener.WithIP(c.ClientIP()))
		if err != nil {
			switch err {
			case shortener.ErrAlreadyExists:
				_ = c.AbortWithError(http.StatusConflict, err)
			default:
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		c.JSON(http.StatusOK, struct {
			Link string `json:"link"`
		}{
			Link: fmt.Sprintf("%s/%s", s.shortHost, link.Code),
		})
	})

	r.GET("/", func(c *gin.Context) {
		code := "default"

		link, err := s.shortener.Lengthen(c, code, shortener.WithIP(c.ClientIP()))
		if err != nil {
			switch err {
			case shortener.ErrNotFound, shortener.ErrExpired, shortener.ErrInvalidIP:
				_ = c.AbortWithError(http.StatusNotFound, err)
			default:
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")

		link, err := s.shortener.Lengthen(c, code, shortener.WithIP(c.ClientIP()))
		if err != nil {
			switch err {
			case shortener.ErrNotFound, shortener.ErrExpired, shortener.ErrInvalidIP:
				_ = c.Error(err)
				c.Data(http.StatusNotFound, "text/html", e404)
			default:
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	})

	return r.Run(s.address)
}
