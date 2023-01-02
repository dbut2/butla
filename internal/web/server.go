package web

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/butla/pkg/database"
	"github.com/dbut2/butla/pkg/datastore"
	"github.com/dbut2/butla/pkg/redis"
	"github.com/dbut2/butla/pkg/shortener"
	"github.com/dbut2/butla/pkg/store"
)

type Server struct {
	address   string
	scheme    string
	hostname  string
	shortener shortener.Shortener
}

func New(config Config) (*Server, error) {
	var s store.Store

	if config.Store.Database.Config != nil {
		db, err := database.New(*config.Store.Database.Config)
		if err != nil {
			return nil, err
		}
		s = db
	}

	if config.Store.Datastore.Config != nil {
		ds, err := datastore.New(*config.Store.Datastore.Config)
		if err != nil {
			return nil, err
		}
		s = ds
	}

	if s == nil {
		s = store.InMem()
	}

	if config.Cache.Redis.Config != nil {
		r, err := redis.New(*config.Cache.Redis.Config)
		if err != nil {
			return nil, err
		}
		s = store.CacheStore{
			Primary: s,
			Cache:   r,
		}
	}

	if config.Host.Scheme == "" {
		config.Host.Scheme = "https"
	}

	return &Server{
		address:   config.Address,
		scheme:    config.Host.Scheme,
		hostname:  config.Host.Hostname,
		shortener: shortener.New(s),
	}, nil
}

func (s *Server) Run() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	t, err := template.ParseFS(pages, "*")
	if err != nil {
		return err
	}
	r.SetHTMLTemplate(t)

	r.GET("/shorten", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"scheme": s.scheme,
		})
	})

	r.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.GET("/500", func(c *gin.Context) {
		c.HTML(http.StatusInternalServerError, "500.html", nil)
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

		u, err := url.Parse(b.Url)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if u.Scheme == "" {
			u.Scheme = "https"
		}

		link, err := s.shortener.Shorten(c, u.String(), shortener.WithExpiry(time.Now().Add(time.Minute*10)), shortener.WithIP(c.ClientIP()))
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
			Link: fmt.Sprintf("%s/%s", s.hostname, link.Code),
		})
	})

	r.GET("/", func(c *gin.Context) {
		code := "default"

		s.lengthen(c, code)
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")

		s.lengthen(c, code)
	})

	return r.Run(s.address)
}

func (s *Server) lengthen(c *gin.Context, code string) {
	link, err := s.shortener.Lengthen(c, code, shortener.WithIP(c.ClientIP()))
	if err != nil {
		_ = c.Error(err)
		switch err {
		case shortener.ErrNotFound, shortener.ErrExpired, shortener.ErrInvalidIP:
			c.HTML(http.StatusNotFound, "404.html", nil)
		default:
			c.HTML(http.StatusInternalServerError, "500.html", nil)
		}
		return
	}

	c.Redirect(http.StatusMovedPermanently, link.Url)
}
