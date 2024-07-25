package grpcapp

import (
	authgrpc "authservice/internal/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
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
	port int,
) App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)

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

	if err := a.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server listening on", "adress", fmt.Sprintf("%v:%d", listener.Addr(), a.port))

	return nil
}

func (a *App) GracefulStop() {
	const op = "grpcapp.GracefulStop"

	a.log.With(slog.String("operation", op)).
		Info("gRPC server gracefully stopping")

	a.gRPCServer.GracefulStop()
}
