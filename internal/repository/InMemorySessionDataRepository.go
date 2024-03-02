package repository

import "github.com/go-webauthn/webauthn/webauthn"

type InMemorySessionDataReopository struct {
	Sessions []webauthn.SessionData
}

func (i *InMemorySessionDataReopository) GetSessionData() (*webauthn.SessionData, error) {
	return &i.Sessions[0], nil
}

func (i *InMemorySessionDataReopository) SaveSessionData(sessionData *webauthn.SessionData) error {
	i.Sessions = append(i.Sessions, *sessionData)
	return nil
}
