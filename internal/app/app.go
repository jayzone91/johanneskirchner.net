package app

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	config     Config
	files      fs.FS
	logger     *slog.Logger
	ipresolver *realip.Service
}

func New(logger *slog.Logger, config Config, files fs.FS) (*App, error) {
	return &App{
		config:     config,
		logger:     logger,
		files:      files,
		upresolver: realip.New(realip.LstXFFIResolver),
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	router, err := a.loadRoutes()
	if err != nil {
		return fmt.Errorf("failed when loading routes: %w", err)
	}

	middlewares := middleware.Chain(a.ipresolver.Middleware(), middleware.Loggin(a.logger), flash.Middleware)

	port := getPort(3000)
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        middlewares(router),
		MaxHeaderBytes: 1 << 20,
	}

	errCh := make(chan error, 1)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to listen and serve: %w", err)
		}

		close(errCh)
	}()

	a.logger.Info("server runnuing", slog.Int("port", port))

	select {
	case <-ctx.Done():
		break
	case err := <-errCh:
		return err
	}

	sCtx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(sCtx)
	return nil
}

func getPort(defaultPort int) int {
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return defaultPort
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return defaultPort
	}
	return port
}
