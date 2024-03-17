package repository

import "fido2server/pkg/webauthn"

type InMemoryUserRepository struct {
	Users map[string]*webauthn.RegisteredUser
}

func (i *InMemoryUserRepository) GetUser(userName string) (*webauthn.RegisteredUser, error) {
	user, ok := i.Users[userName]
	if !ok {
		return nil, webauthn.ErrRegisterUserNotFound
	}
	return user, nil
}

func (i *InMemoryUserRepository) SaveUser(user *webauthn.RegisteredUser) error {
	i.Users[user.UserName] = user
	return nil
}
