package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dbut2/shortener/pkg/datastore"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"

	"github.com/dbut2/shortener/pkg/shortener"
)

type Server struct {
	address   string
	shortHost string
	shortener shortener.Shortener
}

func New(config Config) (*Server, error) {
	ds, err := datastore.NewDatastore(config.Datastore)
	if err != nil {
		return nil, err
	}

	return &Server{
		address:   config.Address,
		shortHost: config.ShortHost,
		shortener: shortener.New(ds),
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
		ctx, span := trace.StartSpan(c, "POST"+c.FullPath())
		defer span.End()
		b := struct {
			Url string `json:"url"`
		}{}
		err := c.BindJSON(&b)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		link, err := s.shortener.Shorten(ctx, b.Url, shortener.WithExpiry(time.Now().Add(time.Minute*10)), shortener.WithIP(c.ClientIP()))
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
		ctx, span := trace.StartSpan(c, "GET"+c.FullPath())
		defer span.End()
		code := "default"

		link, err := s.shortener.Lengthen(ctx, code, shortener.WithIP(c.ClientIP()))
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
		ctx, span := trace.StartSpan(c, "GET"+c.FullPath())
		defer span.End()
		code := c.Param("code")

		link, err := s.shortener.Lengthen(ctx, code, shortener.WithIP(c.ClientIP()))
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
