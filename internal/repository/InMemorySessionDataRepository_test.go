package repository

import (
	"testing"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/go-cmp/cmp"
)

func TestInMemorySessionDatRepository_GetSessionData(t *testing.T) {
	cases := []struct {
		name string

		userName    string
		sessionData map[string]*webauthn.SessionData

		wantSessionData *webauthn.SessionData
		wantErr         bool
	}{
		{
			name: "ShouldInMemorySessionDataRepository_GetSessionDataSuccessfully",

			userName: "test",
			sessionData: map[string]*webauthn.SessionData{
				"test": {
					Challenge: "test",
					UserID:    []byte("test"),
				},
			},

			wantSessionData: &webauthn.SessionData{
				Challenge: "test",
				UserID:    []byte("test"),
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// GIVEN
			sessionDataRepository := &InMemorySessionDataReopository{
				Sessions: c.sessionData,
			}

			// WHEN
			gotSessionData, gotErr := sessionDataRepository.GetSessionData(c.userName)

			// THEN
			if c.wantErr {
				if gotErr == nil {
					t.Errorf("InMemorySessionDataRepository.GetSessionData() expected error, but nil")
				}

				if gotSessionData != nil {
					t.Errorf("InMemorySessionDataRepository.GetSessionData() expected nil, but %v", gotSessionData)
				}
			}

			if diff := cmp.Diff(c.wantSessionData, gotSessionData); diff != "" {
				t.Errorf("InMemorySessionDataRepository.GetSessionData() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
