package grpcapp

import (
	"authservice/internal/domain/auth"
	authgrpc "authservice/internal/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func Constructor(
	log *slog.Logger,
	authService auth.Service,
	port int,
) App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer, authService)
	reflection.Register(gRPCServer)

	return App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Start() error {
	const op = "grpcapp.Start"

	log := a.log.With(slog.String("operation", op))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server listening on", "adress", fmt.Sprintf("%v", listener.Addr()))

	if err := a.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) GracefulStop() {
	const op = "grpcapp.GracefulStop"

	a.log.With(slog.String("operation", op)).
		Info("gRPC server gracefully stopping")

	a.gRPCServer.GracefulStop()
}
