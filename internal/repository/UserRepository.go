package repository

import "github.com/go-webauthn/webauthn/webauthn"

type UserRepository interface {
	GetUser() (webauthn.User, error)
	SaveUser(user webauthn.User) error
}
