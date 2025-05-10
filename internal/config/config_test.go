package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		name string

		// GIVEN
		filepath string

		// THEN
		wantErr    bool
		wantConfig *Config
	}{
		{
			name: "ShouldNewConfigSuccessfully",

			filepath: "../../testdata/config.yaml",

			wantErr: false,
			wantConfig: &Config{
				API: API{
					Port:               8080,
					ShutdownTimeoutSec: 5,
				},
				WebAuthn: WebAuthn{
					RPDisplayName: "testDisplayName",
					RPID:          "testRPID",
					RPOrigin:      "https://test.local",
				},
			},
		},
		{
			name: "ShouldReturnErrorWhenConfigFileIsInvalid",

			filepath: "../../testdata/invalid_config.yaml",

			wantErr: true,
		},
		{
			name: "ShouldReturnErrorWHenCofngiFileIsNotExist",

			filepath: "/path/to/not-exist-file.yaml",

			wantErr: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// WHEN
			gotConfig, gotErr := NewConfig(c.filepath)

			// THEN
			if c.wantErr {
				if gotErr == nil {
					t.Errorf("NewConfig() expected error, but nil")
				}

				if gotConfig != nil {
					t.Errorf("NewConfig() expected nil, but %v", gotConfig)
				}
			}

			if diff := cmp.Diff(c.wantConfig, gotConfig); diff != "" {
				t.Errorf("NewConfig mismatch (-want +got):\n%s", diff)
			}
		})
	}

}
