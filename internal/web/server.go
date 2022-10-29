package web

import (
	"fmt"
	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/shortener-web/pkg/datastore"
	"github.com/dbut2/shortener-web/pkg/shortener"
	"github.com/dbut2/shortener-web/pkg/store"
)

type Server struct {
	address   string
	shortHost string
	shortener shortener.Shortener
}

func New(config Config) (*Server, error) {
	var s store.Store

	if config.Store.Database != nil {
		db, err := database.NewDatabase(*config.Store.Database)
		if err != nil {
			return nil, err
		}
		s = db
	}

	if config.Store.Datastore != nil {
		ds, err := datastore.NewDatastore(*config.Store.Datastore)
		if err != nil {
			return nil, err
		}
		s = ds
	}

	if config.Cache.Redis != nil {
		r, err := redis.NewRedis(*config.Cache.Redis)
		if err != nil {
			return nil, err
		}
		s = store.CacheStore{
			Primary: s,
			Cache:   r,
		}
	}

	return &Server{
		address:   config.Address,
		shortHost: config.ShortHost,
		shortener: shortener.New(s),
	}, nil
}

func (s *Server) Run() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/shorten", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", index)
	})

	r.GET("/404", func(c *gin.Context) {
		c.Data(http.StatusNotFound, "text/html", e404)
	})

	r.GET("/500", func(c *gin.Context) {
		c.Data(http.StatusInternalServerError, "text/html", e500)
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
			_ = c.Error(err)
			switch err {
			case shortener.ErrNotFound, shortener.ErrExpired, shortener.ErrInvalidIP:
				c.Data(http.StatusNotFound, "text/html", e404)
			default:
				c.Data(http.StatusNotFound, "text/html", e500)
			}
			return
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")

		link, err := s.shortener.Lengthen(c, code, shortener.WithIP(c.ClientIP()))
		if err != nil {
			_ = c.Error(err)
			switch err {
			case shortener.ErrNotFound, shortener.ErrExpired, shortener.ErrInvalidIP:
				c.Data(http.StatusNotFound, "text/html", e404)
			default:
				c.Data(http.StatusNotFound, "text/html", e500)
			}
			return
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	})

	return r.Run(s.address)
}
