package web_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dbut2/shortener/internal/web"
	"github.com/dbut2/shortener/pkg/models"
	shortener2 "github.com/dbut2/shortener/pkg/shortener"
	"github.com/dbut2/shortener/pkg/shortener/_mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestNew(t *testing.T) {
	s, err := web.New(web.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestServer_AttachTo(t *testing.T) {
	s := &web.Server{}

	r := gin.Default()
	s.AttachTo(r)
}

func TestServer_routes(t *testing.T) {

	ctrl := gomock.NewController(t)
	shortener := mock_shortener.NewMockShortener(ctrl)

	s := &web.Server{
		Address:   ":8080",
		ShortHost: "localhost",
		Shortener: shortener,
	}
	r := gin.New()
	s.AttachTo(r)

	t.Run("GET", func(t *testing.T) {

		t.Run("shorten", func(t *testing.T) {

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/shorten", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
		})

		t.Run("404", func(t *testing.T) {

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/404", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 404, w.Code)
		})

		t.Run("500", func(t *testing.T) {

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/500", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 500, w.Code)
		})

		t.Run("default", func(t *testing.T) {

			t.Run("success", func(t *testing.T) {

				shortener.EXPECT().Lengthen(gomock.Any(), "default", gomock.Any())

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.ServeHTTP(w, req)

				assert.Equal(t, 301, w.Code)
			})

			t.Run("not found", func(t *testing.T) {

				shortener.EXPECT().Lengthen(gomock.Any(), "default", gomock.Any()).Return(models.Link{}, shortener2.ErrNotFound)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.ServeHTTP(w, req)

				assert.Equal(t, 404, w.Code)
			})

			t.Run("shortener error", func(t *testing.T) {

				shortener.EXPECT().Lengthen(gomock.Any(), "default", gomock.Any()).Return(models.Link{}, shortener2.ErrUnspecified)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.ServeHTTP(w, req)

				assert.Equal(t, 404, w.Code)
			})
		})

		t.Run("code", func(t *testing.T) {

			t.Run("success", func(t *testing.T) {

				shortener.EXPECT().Lengthen(gomock.Any(), "code", gomock.Any())

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/code", nil)
				r.ServeHTTP(w, req)

				assert.Equal(t, 301, w.Code)
			})

			t.Run("not found", func(t *testing.T) {

				shortener.EXPECT().Lengthen(gomock.Any(), "code", gomock.Any()).Return(models.Link{}, shortener2.ErrNotFound)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/code", nil)
				r.ServeHTTP(w, req)

				assert.Equal(t, 404, w.Code)
			})

			t.Run("shortener error", func(t *testing.T) {

				shortener.EXPECT().Lengthen(gomock.Any(), "code", gomock.Any()).Return(models.Link{}, shortener2.ErrUnspecified)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/code", nil)
				r.ServeHTTP(w, req)

				assert.Equal(t, 404, w.Code)
			})
		})
	})

	t.Run("POST", func(t *testing.T) {

		t.Run("shorten", func(t *testing.T) {

			t.Run("success", func(t *testing.T) {

				link := models.Link{
					Code: "somecode",
					Url:  "https://google.com",
				}

				shortener.EXPECT().Shorten(gomock.Any(), link.Url, gomock.Any()).Return(link, nil)

				b, err := json.Marshal(struct {
					Url string `json:"url"`
				}{
					Url: link.Url,
				})
				assert.NoError(t, err)
				body := bytes.NewBuffer(b)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/shorten", body)
				r.ServeHTTP(w, req)

				expected, err := json.Marshal(struct {
					Link string `json:"link"`
				}{
					Link: s.ShortHost + "/" + link.Code,
				})
				assert.NoError(t, err)

				assert.Equal(t, 200, w.Code)
				assert.Equal(t, string(expected), w.Body.String())
			})

			t.Run("bad json", func(t *testing.T) {

				body := bytes.NewBufferString("{bad json")

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/shorten", body)
				r.ServeHTTP(w, req)

				assert.Equal(t, 400, w.Code)
			})

			t.Run("code exists", func(t *testing.T) {

				link := models.Link{
					Code: "somecode",
					Url:  "https://google.com",
				}

				shortener.EXPECT().Shorten(gomock.Any(), link.Url, gomock.Any()).Return(models.Link{}, shortener2.ErrAlreadyExists)

				b, err := json.Marshal(struct {
					Url string `json:"url"`
				}{
					Url: link.Url,
				})
				assert.NoError(t, err)
				body := bytes.NewBuffer(b)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/shorten", body)
				r.ServeHTTP(w, req)

				assert.Equal(t, 409, w.Code)
			})

			t.Run("shortener error", func(t *testing.T) {

				link := models.Link{
					Code: "somecode",
					Url:  "https://google.com",
				}

				shortener.EXPECT().Shorten(gomock.Any(), link.Url, gomock.Any()).Return(models.Link{}, shortener2.ErrUnspecified)

				b, err := json.Marshal(struct {
					Url string `json:"url"`
				}{
					Url: link.Url,
				})
				assert.NoError(t, err)
				body := bytes.NewBuffer(b)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/shorten", body)
				r.ServeHTTP(w, req)

				assert.Equal(t, 500, w.Code)
			})
		})
	})
}
