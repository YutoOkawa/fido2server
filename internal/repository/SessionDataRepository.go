package repository

import "github.com/go-webauthn/webauthn/webauthn"

type SessionDataRepository interface {
	GetSessionData(userName string) (*webauthn.SessionData, error)
	SaveSessionData(sessionData *webauthn.SessionData, userName string) error
}
