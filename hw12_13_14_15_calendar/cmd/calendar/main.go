package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/config"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage/storageconf"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
	config, err := config.NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logg := logger.New(config.Logger.Level)
	storage := storageconf.ChangeStorage(config, logg)
	calendar := app.New(logg, storage)
	httpserver := internalhttp.NewServer(logg, calendar, config.HTTPServer.Host, config.HTTPServer.Port)
	grpcserver := internalgrpc.NewServer(logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	
	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpserver.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}

		if err := grpcserver.Stop(ctx); err != nil {
			logg.Error("failed to stop grpc server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := httpserver.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}()

	go func() {
		defer wg.Done()
		if err := grpcserver.Start(ctx, config.GRPCServer.Network, config.GRPCServer.Address); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}()

	<-ctx.Done()
	wg.Wait()
}
