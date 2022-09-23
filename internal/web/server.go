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
	Address   string
	ShortHost string
	Shortener shortener.Shortener
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
		Address:   config.Address,
		ShortHost: config.ShortHost,
		Shortener: shortener.New(store.Log(s, "main")),
	}, nil
}

func (s *Server) AttachTo(r gin.IRouter) {
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

		link, err := s.Shortener.Shorten(c, b.Url, shortener.WithExpiry(time.Now().Add(time.Minute*10)), shortener.WithIP(c.ClientIP()))
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
			Link: fmt.Sprintf("%s/%s", s.ShortHost, link.Code),
		})
	})

	r.GET("/", func(c *gin.Context) {
		code := "default"

		link, err := s.Shortener.Lengthen(c, code, shortener.WithIP(c.ClientIP()))
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

		link, err := s.Shortener.Lengthen(c, code, shortener.WithIP(c.ClientIP()))
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
}
