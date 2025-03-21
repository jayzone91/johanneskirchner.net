package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type Handler struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(logger *slog.Logger, db *sql.DB) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

func Component(comp templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		comp.Render(r.Context(), w)
	})
}
