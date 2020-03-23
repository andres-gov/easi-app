package server

import (
	"net/http"

	"go.uber.org/zap"

	requestcontext "github.com/cmsgov/easi-app/pkg/context"
)

func traceMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(requestcontext.WithTrace(ctx)))
	})
}

// NewTraceMiddleware returns a handler with a trace ID in context
func NewTraceMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return traceMiddleware(logger, next)
	}
}