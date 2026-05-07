package db

import (
	"github.com/qdrant/go-client/qdrant"
	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/config"
)

func NewVectorStorage(cfg config.Storage) (*qdrant.Client, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: cfg.Host,
		Port: cfg.Port,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
