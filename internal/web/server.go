package web

import (
	"crypto/tls"
	"fmt"
	"net/http"

	pb "github.com/dbut2/shortener/pkg/api/shortener/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	address   string
	shortHost string
	client    pb.ShortServiceClient
}

func New(config Config) (*Server, error) {
	creds := credentials.NewTLS(&tls.Config{})
	if config.Api.Insecure {
		creds = insecure.NewCredentials()
	}

	cc, err := grpc.Dial(config.Api.Host, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client := pb.NewShortServiceClient(cc)

	return &Server{
		address:   config.Address,
		shortHost: config.ShortHost,
		client:    client,
	}, nil
}

func (s *Server) Run() error {
	r := gin.Default()

	r.GET("/shorten", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(index))
	})

	r.GET("/", func(c *gin.Context) {
		resp, err := s.client.Lengthen(c, &pb.LengthenRequest{Code: "default"})
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		long := resp.GetUrl()

		c.Redirect(http.StatusTemporaryRedirect, long)
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")

		resp, err := s.client.Lengthen(c, &pb.LengthenRequest{Code: code})

		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		long := resp.GetUrl()

		c.Redirect(http.StatusMovedPermanently, long)
	})

	r.POST("/shorten", func(c *gin.Context) {
		req := ShortenRequest{}

		err := c.BindJSON(&req)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		resp, err := s.client.Shorten(c, &pb.ShortenRequest{
			Url:  req.Url,
			Code: req.Code,
		})

		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		link := fmt.Sprintf("%s/%s", s.shortHost, resp.GetCode())

		c.JSON(http.StatusOK, ShortenResponse{
			Link: link,
		})
	})

	return r.Run(s.address)
}
