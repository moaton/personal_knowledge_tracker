package ginlogr

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
)

func Logger(l logr.Logger, skipPaths ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(skipPaths); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skipPaths {
			skip[path] = struct{}{}
		}
	}

	return func(ctx *gin.Context) {
		start := time.Now()

		path := ctx.Request.URL.Path

		ctx.Next()

		if _, ok := skip[path]; ok {
			return
		}

		end := time.Now()
		latency := end.Sub(start)

		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors {
				l.Error(e.Unwrap(), "[GIN]")
			}
			return
		}

		l.Info("[GIN]",
			"method", ctx.Request.Method,
			"path", path,
			"status", ctx.Writer.Status(),
			"ip", ctx.ClientIP(),
			"latency", latency,
		)
	}
}
