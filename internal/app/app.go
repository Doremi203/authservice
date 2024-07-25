package app

import (
	grpcapp "authservice/internal/app/grpc"
	"authservice/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	gRPCApp grpcapp.App
}

func Constructor(
	log *slog.Logger,
	port int,
	dbConfig config.DBConfig,
	authConfig config.AuthConfig,
) App {
	grpcApp := grpcapp.Constructor(log, port)

	return App{
		gRPCApp: grpcApp,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error)
	go func() {
		errCh <- a.gRPCApp.Start()
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errCh:
		return err
	case <-stopCh:
		a.gRPCApp.GracefulStop()
		return nil
	}
}
