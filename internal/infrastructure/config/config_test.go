package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/config"
)

func TestLoadFromEnvironment(t *testing.T) {
	t.Setenv("ENV", "production")
	t.Setenv("QDRANT_HOST", "test-host")
	t.Setenv("QDRANT_PORT", "1234")
	t.Setenv("COLLECTION_NAME", "test")

	config, err := config.Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if config.Storage.Host != "test-host" {
		t.Errorf("expected Storage.Host to be 'test-host', got '%s'", config.Storage.Host)
	}

	if config.Storage.Port != 1234 {
		t.Errorf("expected Storage.Port to be 1234, got %d", config.Storage.Port)
	}
}

func TestLoadDevelopmentDotEnv(t *testing.T) {
	clearEnvironment(t, []string{"QDRANT_HOST", "QDRANT_PORT", "COLLECTION_NAME"})

	t.Setenv("ENV", "development")
	tempDir := t.TempDir()
	writeEnvFile(t, tempDir, "QDRANT_HOST=localhost\nQDRANT_PORT=6333")
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get original director %v", err)
	}

	t.Cleanup(func() { _ = os.Chdir(originalDir) })

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	_, err = config.Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}

func TestLoadProdDoesNotReadDotEnv(t *testing.T) {
	clearEnvironment(t, []string{"QDRANT_HOST", "QDRANT_PORT", "COLLECTION_NAME"})
	t.Setenv("ENV", "production")
	tempDir := t.TempDir()

	writeEnvFile(t, tempDir, "")
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get original director %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalDir) })
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	if _, err := config.Load(); err == nil {
		t.Fatalf("expected to fail due to missing environment variables, got no error")
	}
}

func clearEnvironment(t *testing.T, keys []string) {
	t.Helper()

	for _, k := range keys {
		v, exists := os.LookupEnv(k)
		if exists {
			if err := os.Unsetenv(k); err != nil {
				t.Fatalf("failed to unset environment variable %s: %v", k, err)
			}
			k := k
			v := v
			t.Cleanup(func() {
				_ = os.Setenv(k, v)
			})
		}
	}
}

func writeEnvFile(t *testing.T, dir, contents string) {
	t.Helper()

	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte(contents), 0o600); err != nil {
		t.Fatalf("Failed to write .env file: %v", err)
	}
}
