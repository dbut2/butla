package lengthener

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/butla/internal/pages"
	"github.com/dbut2/butla/pkg/shortener"
	"github.com/dbut2/butla/pkg/stores"
)

type Server struct {
	address   string
	shortener shortener.Shortener
}

func New(config *Config) (*Server, error) {
	store, err := stores.New(config.Store)
	if err != nil {
		return nil, err
	}

	return &Server{
		address:   config.Address,
		shortener: shortener.New(store),
	}, nil
}

func (s *Server) Run() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	err := pages.SetDefaults(r)
	if err != nil {
		return err
	}

	r.GET("/", s.lengthen)
	r.GET("/:code", s.lengthen)

	return r.Run(s.address)
}

func (s *Server) lengthen(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		code = "default"
	}

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
