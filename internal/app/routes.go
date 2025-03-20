package app

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/jayzone91/johanneskirchner.net/internal/component"
	"github.com/jayzone91/johanneskirchner.net/internal/handler"
)

func (a *App) loadPages(router *http.ServeMux) {
	h := handler.New(a.logger, a.ipresolver)

	router.Handle("GET /{$}", handler.Component(component.Index()))

	router.HandleFunc("POST /{$}", h.CreateSnippet)
	router.HandleFunc("GET /{$}", h.GetSnippet)

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (a *App) loadStaticFiles() (http.Handler, error) {
	if os.Getenv("BUILD_MODE") == "develop" {
		return http.FileServer(http.Dir("./static")), nil
	}

	static, err := fs.Sub(a.files, "static")
	if err != nil {
		return nil, fmt.Errorf("failed to subdi static: %w", err)
	}

	return http.FileServerFS(static), nil
}

func (a *App) loadRoutes() (http.Handler, error) {
	static, err := a.loadStaticFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to load static files: %w", err)
	}

	// Create a new router
	router := http.NewServeMux()

	// This is the static fileserver.
	router.Handle("GET /static/", http.StripPrefix("/static", static))

	a.loadPages(router)

	return router, nil
}
