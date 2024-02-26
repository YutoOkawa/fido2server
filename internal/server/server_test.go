package server

import (
	"fido2server/internal/config"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// GIVEN
	cfg := config.API{
		Port:               8080,
		ShutdownTimeoutSec: 5,
	}

	expectedPort := ":8080"
	expectedShutdownTimeout := 5 * time.Second

	// WHEN
	server := NewServer(cfg)

	// THEN
	if expectedPort != server.port {
		t.Errorf("expected Port %s, but %s", expectedPort, server.port)
	}

	if expectedShutdownTimeout != server.shutdownTimeout {
		t.Errorf("expected ShutdownTimeout %v, got %v", expectedShutdownTimeout, server.shutdownTimeout)
	}
}
