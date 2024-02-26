package server

import (
	"fido2server/internal/config"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App             *fiber.App
	port            string
	shutdownTimeout time.Duration
}

func NewServer(cfg config.API) *Server {
	app := fiber.New()
	return &Server{
		App:             app,
		port:            fmt.Sprintf(":%d", cfg.Port),
		shutdownTimeout: time.Duration(cfg.ShutdownTimeoutSec) * time.Second,
	}
}

func (s *Server) Start(errCh chan<- error) {
	if err := s.App.Listen(s.port); err != nil {
		errCh <- err
	}
}

func (s *Server) Shutdown() error {
	return s.App.ShutdownWithTimeout(s.shutdownTimeout)
}
