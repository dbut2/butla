package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/butla/pkg/pages"
	"github.com/dbut2/butla/pkg/shortener"
	"github.com/dbut2/butla/pkg/stores"
)

type Server struct {
	address   string
	scheme    string
	hostname  string
	shortener shortener.Shortener
}

func New(config *Config) (*Server, error) {
	s, err := stores.New(config.Store)
	if err != nil {
		return nil, err
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

	err := pages.SetDefaults(r)
	if err != nil {
		return err
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"scheme": s.scheme,
		})
	})

	r.POST("/shorten", s.shorten)

	return r.Run(s.address)
}

func (s *Server) shorten(c *gin.Context) {
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
}
