package bot

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	"github.com/NChitty/archaeologist/cmd/bot/web"
)

//go:embed templates/*.html
var templates embed.FS

func main() {
  logger := slog.New(slog.Default().Handler())

  ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
  defer cancel()

  app := web.New(logger, templates)
	if err := app.Start(ctx); err != nil {
		logger.Error("Failed to start server", slog.Any("error", err))
	}
}
