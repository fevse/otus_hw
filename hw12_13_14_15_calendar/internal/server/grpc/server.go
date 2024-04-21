//go:generate protoc -I ../../../api/ EventService.proto --go_out=./pb/ --go-grpc_out=./pb/
package internalgrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"google.golang.org/grpc"
)

type CalendarServer struct {
	pb.UnimplementedCalendarServer
	App *app.App
	Logger Logger
	Server *grpc.Server
}

type Logger interface {
	Info(string)
	Error(string)
}

func NewServer(logger Logger, app *app.App) *CalendarServer {
	return &CalendarServer{
		App: app,
		Logger: logger,
	}
}

func (s *CalendarServer) Start (ctx context.Context,network string, address string) error {
		s.Logger.Info("server is running: " + address)
		listen, err := net.Listen(network, address)
		if err != nil {
			s.Logger.Error("failed to start grpcserver" + err.Error())
		}
		
		s.Server = grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor(s.Logger)))
		pb.RegisterCalendarServer(s.Server, s)
		
		return s.Server.Serve(listen)
}

func (s *CalendarServer) Stop(ctx context.Context) error {
	s.Server.GracefulStop()
	return nil
}

func loggingInterceptor(logg Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		logg.Info(info.FullMethod + " " + fmt.Sprintf("%v", req))
		return handler(ctx, req)
	}
}
