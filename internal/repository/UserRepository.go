package repository

import "fido2server/pkg/webauthn"

type UserRepository interface {
	GetUser(userName string) (*webauthn.RegisteredUser, error)
	SaveUser(user *webauthn.RegisteredUser) error
}
