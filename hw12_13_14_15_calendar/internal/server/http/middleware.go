package internalhttp

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler { //nolint:unused
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.Logger.Info(fmt.Sprintf("%s %s %s %s %v %v %s", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, http.StatusOK, time.Since(start), r.UserAgent()))
		// TODO
	})
}
