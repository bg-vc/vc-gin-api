package pkg

import (
	"errors"
	"testing"
	"time"
)

func TestTokenRoundTrip(t *testing.T) {
	manager := NewTokenManager("test-secret-that-is-at-least-32-characters", 30*time.Minute)
	token, expiresAt, err := manager.Sign(7, "alice")
	if err != nil {
		t.Fatal(err)
	}

	context, err := manager.ParseAuthorization("Bearer " + token)
	if err != nil {
		t.Fatal(err)
	}
	if context.AccountID != 7 || context.Username != "alice" {
		t.Fatalf("unexpected context: %+v", context)
	}
	if !context.ExpiresAt.Equal(expiresAt) {
		t.Fatalf("unexpected expiry: got %s want %s", context.ExpiresAt, expiresAt)
	}
}

func TestTokenRejectsInvalidAuthorization(t *testing.T) {
	manager := NewTokenManager("test-secret-that-is-at-least-32-characters", 30*time.Minute)
	_, err := manager.ParseAuthorization("not-a-bearer-token")
	if !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("expected ErrInvalidToken, got %v", err)
	}
}

func TestTokenRejectsExpiredToken(t *testing.T) {
	manager := NewTokenManager("test-secret-that-is-at-least-32-characters", time.Minute)
	manager.now = func() time.Time { return time.Now().Add(-2 * time.Minute) }
	token, _, err := manager.Sign(7, "alice")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := manager.ParseAuthorization("Bearer " + token); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("expected expired token to fail, got %v", err)
	}
}
