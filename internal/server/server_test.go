package server

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// GIVEN
	port := 8080
	shutdownTimeout := 5

	// WHEN
	server := NewServer(port, shutdownTimeout)

	// THEN
	if server.port != ":8080" {
		t.Errorf("expected :%d to be :8080, got %s", port, server.port)
	}

	if server.shutdownTimeout != time.Duration(shutdownTimeout)*time.Second {
		t.Errorf("expected shutdownTimeout to be %v, got %d", time.Duration(server.shutdownTimeout)*time.Second, server.shutdownTimeout)
	}
}
