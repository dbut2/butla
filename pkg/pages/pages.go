package pages

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed *.html
var pages embed.FS

func SetDefaults(r *gin.Engine) error {
	t, err := template.ParseFS(pages, "*.html")
	if err != nil {
		return err
	}
	r.SetHTMLTemplate(t)
	r.NoRoute(statusFunc(http.StatusNotFound))
	r.NoMethod(statusFunc(http.StatusNotFound))
	return nil
}

func statusFunc(status int) func(*gin.Context) {
	return func(c *gin.Context) {
		if _, err := pages.Open(fmt.Sprintf("%d.html", status)); err != nil {
			status = http.StatusInternalServerError
		}
		c.HTML(status, fmt.Sprintf("%d.html", status), nil)
	}
}
