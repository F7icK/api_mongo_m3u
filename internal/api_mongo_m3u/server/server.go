package server

import (
	"context"
	"net/http"

	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/server/handlers"
	"github.com/F7icK/api_mongo_m3u/internal/api_mongo_m3u/types/config"
	"github.com/F7icK/api_mongo_m3u/pkg/logger"
	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handlers *handlers.Handlers) *Server {

	router := NewRouter(handlers)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.HTTP.Port,
			Handler: corsOpts.Handler(router),
		},
	}
}

func (s *Server) Run() error {
	logger.LogInfo("Restart server dev")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
