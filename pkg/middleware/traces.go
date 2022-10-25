package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := trace.StartSpan(c, c.FullPath())
		defer span.End()
		ctx.(*gin.Context).Next()
	}
}
