package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct { // TODO
	Server *http.Server
	Logger Logger
}

type Logger interface { // TODO
	Info(string)
	Error(string)
}

type Application interface { // TODO
}

func NewServer(logger Logger, _ Application, host, port string) *Server {
	return &Server{
		Server: &http.Server{
			Addr:              net.JoinHostPort(host, port),
			ReadHeaderTimeout: 2 * time.Second,
		},
		Logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, user!"))
	})
	s.Server.Handler = s.loggingMiddleware(mux)
	s.Logger.Info("server is running: " + s.Server.Addr)
	if err := s.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server err: %w", err)
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return s.Server.Shutdown(ctx)
}

// TODO
