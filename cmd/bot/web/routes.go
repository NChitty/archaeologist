package web

import (
	"html/template"
	"net/http"

	"github.com/NChitty/archaeologist/cmd/bot/web/handler"
)

func (app *App) loadRoutes() {
	tmpl := template.Must(template.New("").ParseFS(app.templates, "templates/*"))
	handler := handler.New(app.logger, tmpl, app.artifactsClient)
	files := http.FileServer(http.Dir("./static"))

	app.router.Handle("GET /static/", http.StripPrefix("/static", files))

	app.router.Handle("GET /login", http.HandlerFunc(handler.LoginPage))

	app.router.Handle("GET /welcome", http.HandlerFunc(handler.Welcome))

	app.router.Handle("POST /login", http.HandlerFunc(handler.Login))
}
