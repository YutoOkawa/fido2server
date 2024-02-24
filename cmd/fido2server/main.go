package main

import (
	"context"
	"fido2server/internal/server"
	"os/signal"
	"syscall"
)

func main() {
	server := server.NewServer(8080, 5)
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxCancel()

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
