package repository

import (
	"github.com/go-webauthn/webauthn/webauthn"
)

type InMemoryUserRepository struct {
	Users []webauthn.User
}

func (i *InMemoryUserRepository) SaveUser(user webauthn.User) error {
	i.Users = append(i.Users, user)
	return nil
}
