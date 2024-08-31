package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/isutare412/goasis/pkg/oapi"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg Config) *Server {
	r := mux.NewRouter()
	handler := oapi.HandlerWithOptions(
		&handler{},
		oapi.GorillaServerOptions{
			BaseRouter:       r,
			ErrorHandlerFunc: responseError,
		})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}

	return &Server{
		server: server,
	}
}

func (s *Server) Run() <-chan error {
	errs := make(chan error, 1)
	go func() {
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errs <- err
			return
		}
	}()

	return errs
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutting down http server: %w", err)
	}
	return nil
}
