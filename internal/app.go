package app

import (
	"context"
	"errors"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/NChitty/archaeologist/handler"
	"github.com/promiseofcake/artifactsmmo-go-client/client"
)

type App struct {
	logger          *slog.Logger
	router          *http.ServeMux
	templates       fs.FS
	artifactsClient *client.ClientWithResponses
}

func New(logger *slog.Logger, templates fs.FS) *App {
	router := http.NewServeMux()

	app := &App{
		logger:    logger,
		router:    router,
		templates: templates,
	}

	return app
}

func (app *App) Start(ctx context.Context) error {
	tmpl := template.Must(template.New("").ParseFS(app.templates, "templates/*"))
	client, err := client.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		app.logger.Error("Could not create artifacts client")
		return err
	}

	app.artifactsClient = client

	handler := handler.New(app.logger, tmpl, client)
	files := http.FileServer(http.Dir("./static"))

	app.router.Handle("GET /static/", http.StripPrefix("/static", files))

	app.router.Handle("GET /login", http.HandlerFunc(handler.LoginPage))

	app.router.Handle("GET /welcome", http.HandlerFunc(handler.Welcome))

	app.router.Handle("POST /login", http.HandlerFunc(handler.Login))

	server := http.Server{
		Addr:    ":8080",
		Handler: app.router,
	}

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.logger.Error("Failed to listen and serve", slog.Any("error", err))
		}
		close(done)
	}()

	app.logger.Info("Server listening", slog.String("addr", ":8080"))

	select {
	case <-done:
		break
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		server.Shutdown(ctx)
		cancel()
	}

	return nil
}
