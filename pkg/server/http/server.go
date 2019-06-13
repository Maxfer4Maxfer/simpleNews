package server

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"simpleNews/pkg/config"
	"simpleNews/pkg/service"
)

// Server represents HTTP Server
type Server struct {
	server http.Server
}

// New creates an instance of the Server
func New(cfg config.Config, service *service.Service) *Server {
	s := &Server{}

	m := NewHandler(service)

	s.server = http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}

// Run starts HTTP Server
func (s *Server) Run() {
	log.Info("HTTP server started ", s.server.Addr)
	s.server.ListenAndServe()
}
