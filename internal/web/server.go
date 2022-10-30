package web

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/datastore"
	"github.com/dbut2/shortener-web/pkg/redis"
	"github.com/dbut2/shortener-web/pkg/shortener"
	"github.com/dbut2/shortener-web/pkg/store"
)

type Server struct {
	address   string
	scheme    string
	host      string
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

	if s == nil {
		s = store.InMem()
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

	if config.ShortHost.Scheme == "" {
		config.ShortHost.Scheme = "https"
	}

	return &Server{
		address:   config.Address,
		scheme:    config.ShortHost.Scheme,
		host:      config.ShortHost.URL,
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
			Link: fmt.Sprintf("%s/%s", s.host, link.Code),
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
