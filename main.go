package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	app "github.com/NChitty/archaeologist/internal"
)

//go:embed templates/*.html
var templates embed.FS

func main() {
  logger := slog.New(slog.Default().Handler())

  ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
  defer cancel()

  app := app.New(logger, templates)
	if err := app.Start(ctx); err != nil {
		logger.Error("Failed to start server", slog.Any("error", err))
	}
}
