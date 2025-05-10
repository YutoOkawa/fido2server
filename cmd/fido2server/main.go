package main

import (
	"context"
	"fido2server/internal/config"
	"fido2server/internal/handler"
	"fido2server/internal/repository"
	"fido2server/internal/server"
	"fido2server/internal/service"
	webauthnlib "fido2server/pkg/webauthn"
	"os/signal"
	"syscall"
)

// TODO: マウント先に変更(rootで実行できる状態になっている)
var defaultFilePath = "testdata/config.yaml"

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxCancel()

	cfg, err := config.NewConfig(defaultFilePath)
	if err != nil {
		return
	}

	server := server.NewServer(cfg.API)

	webauthn, err := webauthnlib.NewWebAuthn(cfg.WebAuthn)
	if err != nil {
		return
	}

	// initialize repository
	inMemoryUserRepositoryImpl := repository.NewInMemoryUserRepository()
	inMemorySessionRepositoryImpl := repository.NewInMemorySessionDataReopository()

	// initialize service
	registerOptionsService := service.RegisterOptionsService{
		UserRepository:        inMemoryUserRepositoryImpl,
		SessionDataRepository: inMemorySessionRepositoryImpl,
		WebAuthn:              webauthn,
	}
	registerResultService := service.RegisterResultService{
		UserRepository:    inMemoryUserRepositoryImpl,
		SessionRepository: inMemorySessionRepositoryImpl,
		WebAuthn:          webauthn,
	}

	// initialize handler
	server.App.Get("/healthz", handler.HealthzHandler)
	v1 := server.App.Group("v1")
	{
		register := v1.Group("/register")
		{
			register.Post("/options", handler.RegisterOptionsHandler(registerOptionsService))
			register.Post("/result", handler.RegisterResultHandler(registerResultService))
		}
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
