package pkg

import (
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Setenv("VC_JWT_SECRET", strings.Repeat("s", 32))
	t.Setenv("VC_DB_PASSWORD", "local-test-password")
	t.Setenv("VC_REDIS_ENABLED", "true")
	t.Setenv("VC_REDIS_DB", "2")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.Redis.Enabled || cfg.Redis.DB != 2 {
		t.Fatalf("unexpected Redis config: %+v", cfg.Redis)
	}
}

func TestLoadConfigRequiresSecrets(t *testing.T) {
	t.Setenv("VC_JWT_SECRET", "short")
	t.Setenv("VC_DB_PASSWORD", "")
	if _, err := LoadConfig(); err == nil {
		t.Fatal("expected invalid local configuration to fail")
	}
}
