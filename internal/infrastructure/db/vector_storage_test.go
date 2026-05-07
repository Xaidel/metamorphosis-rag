package db_test

import (
	"testing"

	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/config"
	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/db"
)

func TestNewVectorStorage(t *testing.T) {
	cfg := config.Storage{
		Host: "localhost",
		Port: 6333,
	}
	client, err := db.NewVectorStorage(cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("expected a client instance, got nil")
	}
}
