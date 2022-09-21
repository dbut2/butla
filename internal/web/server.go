package web

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/dbut2/shortener-web/pkg/ptr"
	pb "github.com/dbut2/shortener/pkg/api/shortener/v1alpha1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		c.Data(http.StatusOK, "text/html", index)
	})

	r.POST("/shorten", func(c *gin.Context) {
		b := struct {
			Url string `json:"url"`
		}{}

		err := c.BindJSON(&b)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		resp, err := s.client.Shorten(c, &pb.ShortenRequest{
			Url:    b.Url,
			Expiry: timestamppb.New(time.Now().Add(time.Minute * 10)),
			Ip:     ptr.To(c.ClientIP()),
		})
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, struct {
			Link string `json:"link"`
		}{
			Link: fmt.Sprintf("%s/%s", s.shortHost, resp.GetCode()),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "default")
	})

	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		ip := c.ClientIP()

		resp, err := s.client.Lengthen(c, &pb.LengthenRequest{
			Code: code,
			Ip:   ip,
		})

		if err != nil {
			s := status.Convert(err)
			if s.Code() == codes.NotFound {
				_ = c.AbortWithError(http.StatusNotFound, err)
				return
			}

			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		long := resp.GetUrl()
		c.Redirect(http.StatusMovedPermanently, long)
	})

	return r.Run(s.address)
}
