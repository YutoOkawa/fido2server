package repository

import "github.com/go-webauthn/webauthn/webauthn"

type UserRepository interface {
	SaveUser(user webauthn.User) error
}
