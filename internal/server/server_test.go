package server

import "testing"

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

	if server.shutdownTimeout != 5 {
		t.Errorf("expected shutdownTimeout to be %d, got %d", shutdownTimeout, server.shutdownTimeout)
	}
}
