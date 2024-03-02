package repository

import "github.com/go-webauthn/webauthn/webauthn"

type SessionDataRepository interface {
	GetSessionData() (*webauthn.SessionData, error)
	SaveSessionData(sessionData *webauthn.SessionData) error
}
