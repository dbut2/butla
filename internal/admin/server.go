package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dbut2/shortener-web/pkg/database"
	"github.com/dbut2/shortener-web/pkg/datastore"
	"github.com/dbut2/shortener-web/pkg/shortener"
	"github.com/dbut2/shortener-web/pkg/store"
)

type Server struct {
	address   string
	shortener shortener.AdminShortener
}

func New(config Config) (*Server, error) {
	var s store.AdminStore

	if config.Store.Database.C != nil {
		db, err := database.New(*config.Store.Database.C)
		if err != nil {
			return nil, err
		}
		s = db
	}

	if config.Store.Datastore.C != nil {
		ds, err := datastore.New(*config.Store.Datastore.C)
		if err != nil {
			return nil, err
		}
		s = ds
	}

	return &Server{
		address:   config.Address,
		shortener: shortener.NewAdmin(s),
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

	r.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.GET("/500", func(c *gin.Context) {
		c.HTML(http.StatusInternalServerError, "500.html", nil)
	})

	r.GET("/list", func(c *gin.Context) {
		links, err := s.shortener.LengthenAll(c)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.String(http.StatusOK, fmt.Sprint(links))
	})

	r.GET("/delete/:code", func(c *gin.Context) {
		code := c.Param("code")

		err := s.shortener.Delete(c, code)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	})

	return r.Run(s.address)
}
