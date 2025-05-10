package webauthn

import (
	"crypto/rand"

	webauthnlib "github.com/go-webauthn/webauthn/webauthn"
)

type RegisteredUser struct {
	ID          []byte
	UserName    string
	DisplayName string
	Icon        string
	Credentials []webauthnlib.Credential
}

func NewRegisteredUser(userName, displayName, icon string) (*RegisteredUser, error) {
	id := make([]byte, 64)
	_, err := rand.Read(id)
	if err != nil {
		return nil, err
	}
	return &RegisteredUser{
		ID:          id,
		UserName:    userName,
		DisplayName: displayName,
		Icon:        icon,
	}, nil
}

func (n *RegisteredUser) WebAuthnID() []byte {
	return n.ID
}

func (n *RegisteredUser) WebAuthnName() string {
	return n.UserName
}

func (n *RegisteredUser) WebAuthnDisplayName() string {
	return n.DisplayName
}

func (n *RegisteredUser) WebAuthnIcon() string {
	return n.Icon
}

func (n *RegisteredUser) WebAuthnCredentials() []webauthnlib.Credential {
	return n.Credentials
}

func (r *RegisteredUser) AddCredential(credential webauthnlib.Credential) {
	r.Credentials = append(r.Credentials, credential)
}
