package repository

import (
	webauthnlib "fido2server/pkg/webauthn"

	"github.com/go-webauthn/webauthn/webauthn"
)

type InMemorySessionDataReopository struct {
	Sessions map[string]*webauthn.SessionData
}

func (i *InMemorySessionDataReopository) GetSessionData(userName string) (*webauthn.SessionData, error) {
	session, ok := i.Sessions[userName]
	if !ok {
		return nil, webauthnlib.ErrSessionDataNotFound
	}
	return session, nil
}

func (i *InMemorySessionDataReopository) SaveSessionData(sessionData *webauthn.SessionData, userName string) error {
	i.Sessions[userName] = sessionData
	return nil
}
