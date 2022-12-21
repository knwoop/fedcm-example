package http

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(port int) *Server {
	mux := http.NewServeMux()
	mux.Handle("/users", &listUserHandler{})
	mux.Handle("/user", &getUserHandler{})
	mux.Handle("/me", &getMeHandler{})

	mux.Handle("/auth/signin", &signinHandler{})

	mux.Handle("/.well-known/web-identity", &getWellKnownFile{})
	mux.Handle("/config.json", &getConfigFile{})

	mux.Handle("/accounts", &accountsHandler{})
	mux.Handle("/metadata", &metadataHandler{})
	mux.Handle("/assertion", &assertionHandler{})

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}
