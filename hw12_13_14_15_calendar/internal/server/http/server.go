package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/app"
)

type Server struct {
	Server *http.Server
	Logger Logger
	App *app.App
}

type Logger interface {
	Info(string)
	Error(string)
}

func NewServer(logger Logger, app *app.App, host, port string) *Server {
	return &Server{
		Server: &http.Server{
			Addr:              net.JoinHostPort(host, port),
			ReadHeaderTimeout: 2 * time.Second,
		},
		App: app,
		Logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()

	mux.Handle("GET /", s.index())
	mux.Handle("POST /create/", s.Create())
	mux.Handle("POST /update/", s.Update())
	mux.Handle("POST /delete/", s.Delete())
	mux.Handle("GET /list/", s.Show())
	mux.Handle("GET /day/", s.ShowEventDay())
	mux.Handle("GET /week/", s.ShowEventWeek())
	mux.Handle("GET /month/", s.ShowEventMonth())

	s.Server.Handler = s.loggingMiddleware(mux)
	s.Logger.Info("server is running: " + s.Server.Addr)
	if err := s.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server err: %w", err)
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
