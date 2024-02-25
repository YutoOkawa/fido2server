package main

import (
	"context"
	"fido2server/internal/handler"
	"fido2server/internal/server"
	"fido2server/internal/service"
	"os/signal"
	"syscall"
)

func main() {
	server := server.NewServer(8080, 5)
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxCancel()

	// initialize service
	registerService := service.RegisterService{}

	// initialize handler
	server.App.Get("/healthz", handler.HealthzHandler)
	v1 := server.App.Group("v1")
	{
		v1.Post("/register", handler.RegisterHandler(registerService))
	}

	srvErrCh := make(chan error, 1)
	go server.Start(srvErrCh)

	for {
		select {
		case err := <-srvErrCh:
			panic(err)
		case <-ctx.Done():
			if err := server.Shutdown(); err != nil {
				panic(err)
			}
			return
		}
	}
}
