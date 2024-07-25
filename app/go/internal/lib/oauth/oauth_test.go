package oauth_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"mickamy.com/playground/internal/lib/oauth"
)

func TestValidateToken(t *testing.T) {
	t.Parallel()
	t.Skip("skipping for token is expired. assign another one on testing this func.")

	// arrange
	tokenStr := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImMzYWJlNDEzYjIyNjhhZTk3NjQ1OGM4MmMxNTE3OTU0N2U5NzUyN2UiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIxMDcxOTM1NTAyOC0xYWVvOGI2dnJmaGFlZ29saHUxcGhnZjFjNTYxcjFjMC5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsImF1ZCI6IjEwNzE5MzU1MDI4LTFhZW84YjZ2cmZoYWVnb2xodTFwaGdmMWM1NjFyMWMwLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTA2MjkzMTgyMjg4NzI0NzY4NDIyIiwiZW1haWwiOiJ0Lm1pa2FtaTE5OTJAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJ3OFNkTFN0clNjSzBzZWJyOVVDRTFnIiwibmFtZSI6IlRldHN1cm8gTWlrYW1pIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0lLb0d0bHgya205cW5jMHQwT3dya0w3THVjTlY4RDB1bDFkMjNrMWVTdmprSnNBUnZzPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IlRldHN1cm8iLCJmYW1pbHlfbmFtZSI6Ik1pa2FtaSIsImlhdCI6MTcxODI1NTk4NywiZXhwIjoxNzE4MjU5NTg3fQ.gMFjTW057zICc1TnVDWw9fKl-QaiIyb2lFh2ep6CwTSJSJVaRQrq9IxcHpm5P-ng_iF-fOmbj-ZIFANFm5mHYLlx82v0pgE56_WbCBkecL1NCCfS6Zq_lEasiP8a5eT3RQuHxIKUOe5W9AgA5UBv7Xu3OjSJoJUMHCn-yhcD6iw3U5Svwg6dX4OUDkiEnAP8W2x31sg1bUOx6aFUnCTMeB9iXK8uZUmhQQRUB98zqPtEFwWNDiKMkE1sIvaiZtnXeJdYmppxPEZG3AGrwg4dzL4WBmUP7C376PKTgYJYFkkE1VzahmiXzkrFihl8gJ9bGxfveFMraryp5ZzwPgmW_g"

	// act
	sut := oauth.New()
	payload, err := sut.ValidateToken(context.Background(), "google", tokenStr)

	t.Logf("payload: %#v", payload)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, payload.Provider)
	assert.NotEmpty(t, payload.UID)
	assert.NotEmpty(t, payload.Name)
	assert.NotEmpty(t, payload.Email)
	assert.NotEmpty(t, payload.Picture)
}
