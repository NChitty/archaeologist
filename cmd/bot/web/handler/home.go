package handler

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/promiseofcake/artifactsmmo-go-client/client"
)

type Archaeologist struct {
	logger *slog.Logger
	tmpl   *template.Template
	client *client.ClientWithResponses
}

func New(
	logger *slog.Logger, tmpl *template.Template, client *client.ClientWithResponses,
) *Archaeologist {
	return &Archaeologist{
		logger: logger,
		tmpl:   tmpl,
		client: client,
	}
}

func (h *Archaeologist) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	h.tmpl.ExecuteTemplate(w, "login.html", nil)
}

func (h *Archaeologist) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	h.tmpl.ExecuteTemplate(w, "welcome.html", nil)
}

func (h *Archaeologist) Login(w http.ResponseWriter, r *http.Request) {
	_, err := h.client.GenerateTokenTokenPostWithResponse(
		context.TODO(),
		client.NewBasicAuthorizationRequestFunc(r.FormValue("email"), r.FormValue("password")),
	)
	if err != nil {
		h.logger.Error("Invalid login attempt", slog.String("error", err.Error()))
	}
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
