package server

import (
	"log"
	"net/http"
	"time"

	"github.com/demidenkoalex/sprint6-final/internal/handlers"
)

type Server struct {
	log    *log.Logger
	server *http.Server
}

func (s *Server) Start() error {
	s.log.Printf("listening on %s", s.server.Addr)

	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func CreateServer(log *log.Logger) *Server {
	router := handlers.NewRouter()
	serv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	server := &Server{
		log:    log,
		server: serv,
	}
	return server
}
