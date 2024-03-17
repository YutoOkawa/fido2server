package webauthn

import (
	"fido2server/internal/config"

	"github.com/go-webauthn/webauthn/webauthn"
)

func NewWebAuthn(cfg config.WebAuthn) (*webauthn.WebAuthn, error) {
	wconfig := &webauthn.Config{
		RPDisplayName: cfg.RPDisplayName,
		RPID:          cfg.RPID,
		RPOrigin:      cfg.RPOrigin,
	}
	webauthn, err := webauthn.New(wconfig)
	if err != nil {
		return nil, err
	}

	return webauthn, nil
}
