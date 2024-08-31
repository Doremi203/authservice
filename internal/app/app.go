package app

import (
	grpcapp "authservice/internal/app/grpc"
	"authservice/internal/domain/auth"
	"authservice/internal/domain/password"
	"authservice/internal/domain/token"
	"authservice/internal/domain/user"
	"authservice/pkg/postgres"
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
	dbConfig postgres.Config,
	tokenConfig token.Config,
) App {
	timeProvider := utils.NewDefaultTimeProvider()
	tokenService := token.NewJWTService(tokenConfig, timeProvider)
	db := postgres.MustNew(dbConfig)
	userRepository := user.NewPostgresRepository(db)
	hashProvider := password.NewBCryptProvider()
	authService := auth.NewDefaultService(tokenService, userRepository, hashProvider)

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
