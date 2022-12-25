package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/knwoop/fedcm-example/idp/db"
)

type Server struct {
	server *http.Server
	db     *db.DB
}

func NewServer(port int) *Server {
	db := db.NewDB()
	s := &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
		db: db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/me", s.getMeHandler)
	mux.HandleFunc("/auth/signin", s.Signin)
	mux.HandleFunc("/auth/idtoken", s.SigninWithIDToken)

	mux.HandleFunc("/.well-known/web-identity", s.GetWellKnownFileHandler)
	mux.HandleFunc("/config.json", s.GetConfigFileHandler)

	mux.HandleFunc("/fedcm/accounts_endpoint", s.AccountsHandler)
	mux.HandleFunc("/fedcm/client_metadata_endpoint", s.MetadataHandler)
	mux.HandleFunc("/fedcm/id_assertion_endpoint", s.AssertionHandler)

	s.server.Handler = mux

	return s
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
