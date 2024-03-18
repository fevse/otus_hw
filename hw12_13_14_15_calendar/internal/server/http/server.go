package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Server struct { // TODO
	Server *http.Server
	Logger Logger	
	Handler http.Handler
}

type Logger interface { // TODO
	Info(string)
	Error(string)
}

type Application interface { // TODO
}

func NewServer(logger Logger, app Application, host, port string) *Server {
	return &Server{
		Server: &http.Server{
			Addr: net.JoinHostPort(host, port),
		},
		Logger: logger,

	}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, user!"))
	})
	s.Server.Handler = s.loggingMiddleware(mux)
	s.Logger.Info("Server is running")
	if err := s.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server err: %v", err)
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return s.Server.Shutdown(ctx)
}

// TODO
