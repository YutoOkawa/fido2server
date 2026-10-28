package webauthn

import (
	"fido2server/internal/config"
	"testing"
)

func TestNewWebAuthn(t *testing.T) {
	cases := []struct {
		name string

		// GIVEN
		cfg config.WebAuthn

		// THEN
		wantErr bool
	}{
		{
			name: "ShouldNewWebAuthnSuccessfully",

			cfg: config.WebAuthn{
				RPDisplayName: "testDisplayName",
				RPID:          "testRPID",
				RPOrigin:      "https://test.local",
			},
		},
		{
			name: "ShouldReturnErrorWhenRPDisplayNameIsEmpty",

			cfg: config.WebAuthn{
				RPDisplayName: "",
				RPID:          "testRPID",
				RPOrigin:      "https://test.local",
			},

			wantErr: true,
		},
		{
			name: "ShouldReturnErrorWhenRPIDIsEmpty",

			cfg: config.WebAuthn{
				RPDisplayName: "testDisplayName",
				RPID:          "",
				RPOrigin:      "https://test.local",
			},

			wantErr: true,
		},
		{
			name: "ShouldReturnErrorWhenRPOriginIsEmpty",
			cfg: config.WebAuthn{
				RPDisplayName: "testDisplayName",
				RPID:          "testRPID",
				RPOrigin:      "",
			},

			wantErr: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// WHEN
			_, gotErr := NewWebAuthn(c.cfg)

			// THEN
			if c.wantErr {
				if gotErr == nil {
					t.Errorf("NewWebAuthn() expected error, but nil")
				}
			} else {
				if gotErr != nil {
					t.Errorf("NewWebAuthn() unexpected error: %v", gotErr)
				}
			}
		})
	}
}
