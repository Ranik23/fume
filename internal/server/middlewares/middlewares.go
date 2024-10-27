package middlewares

import (
	"log/slog"
	"net/http"
	"os"
	"time"
	"github.com/lmittmann/tint"
)

func TimeMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
        h.ServeHTTP(w, r)
		duration := time.Since(start)

		handl := tint.NewHandler(os.Stdout, &tint.Options{})
		logger := slog.New(handl)
		logger.Info("Request completed", "method", r.Method, "url", r.URL.Path, "duration", duration.String())
    })
}