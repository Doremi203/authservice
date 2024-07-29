package main

import (
	"authservice/internal/app"
	"authservice/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting server",
		slog.String("Env", cfg.Env),
		slog.Int("Port", cfg.Grpc.Port))

	application := app.Constructor(
		log,
		cfg.Grpc.Port,
		cfg.DB,
		cfg.Token,
	)

	application.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelError,
			}),
		)
	}

	return log
}
