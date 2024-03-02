package repository

import (
	"fido2server/pkg/webauthn"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInMemoryUserRepository_GetUser(t *testing.T) {
	cases := []struct {
		name string

		userName       string
		displayName    string
		registeredUser map[string]*webauthn.RegisteredUser

		wantErr            bool
		wantRegisteredUser *webauthn.RegisteredUser
	}{
		{
			name: "ShouldInMemoryUserRepository_GetUserSuccessfully",

			userName:    "test",
			displayName: "test",
			registeredUser: map[string]*webauthn.RegisteredUser{
				"test": {
					UserName:    "test",
					DisplayName: "test",
				},
			},

			wantRegisteredUser: &webauthn.RegisteredUser{
				UserName:    "test",
				DisplayName: "test",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// GIVEN
			userRepository := &InMemoryUserRepository{
				Users: c.registeredUser,
			}

			// WHEN
			gotRegisteredUser, gotErr := userRepository.GetUser(c.userName, c.displayName)

			// THEN
			if c.wantErr {
				if gotErr == nil {
					t.Errorf("InMemoryUserRepository.GetUser() expected error, but nil")
				}

				if gotRegisteredUser != nil {
					t.Errorf("InMemoryUserRepository.GetUser() expected nil, but %v", gotRegisteredUser)
				}
			}

			if diff := cmp.Diff(c.wantRegisteredUser, gotRegisteredUser); diff != "" {
				t.Errorf("InMemoryUserRepository.GetUser() diff: (-want +got)\n%s", diff)
			}
		})
	}
}
