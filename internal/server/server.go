package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app             *fiber.App
	port            string
	shutdownTimeout time.Duration
}

func NewServer(port, shutdownTimeout int) *Server {
	app := fiber.New()
	return &Server{
		app:             app,
		port:            fmt.Sprintf(":%d", port),
		shutdownTimeout: time.Duration(shutdownTimeout) * time.Second,
	}
}

func (s *Server) Start(errCh chan<- error) {
	if err := s.app.Listen(s.port); err != nil {
		errCh <- err
	}
}

func (s *Server) Shutdown() error {
	return s.app.ShutdownWithTimeout(s.shutdownTimeout)
}
