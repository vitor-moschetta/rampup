package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-core/pkg/web"
)

func LoggingMiddleware() web.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()

			log.Info(ctx, fmt.Sprintf("Request received: %s %s", r.Method, r.URL.Path))
			log.Info(ctx, fmt.Sprintf("Request headers: %s", r.Header))
			log.Info(ctx, fmt.Sprintf("Request body: %s", r.Body))
			log.Info(ctx, fmt.Sprintf("Request query: %s", r.URL.Query()))

			next(w, r)

			log.Info(ctx, fmt.Sprintf("Request processed in %s", time.Since(start)))
		}
	}
}
