package app

import (
	grpcapp "authservice/internal/app/grpc"
	"authservice/internal/config"
	"authservice/internal/domain/services/auth"
	"authservice/internal/domain/services/token"
	"authservice/internal/domain/services/user"
	"authservice/pkg/utils"
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
	tokenConfig token.Config,
) App {
	timeProvider := utils.NewDefaultTimeProvider()
	tokenService := token.NewJWTService(tokenConfig, timeProvider)
	userRepository := user.NewPostgresRepository()
	authService := auth.NewDefaultService(tokenService, userRepository)

	grpcApp := grpcapp.Constructor(log, authService, port)

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
