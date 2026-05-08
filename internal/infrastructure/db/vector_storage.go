package db

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/qdrant/go-client/qdrant"
	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/config"
)

func NewVectorStorage(cfg config.Storage) (*qdrant.Client, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: cfg.Host,
		Port: cfg.Port,
	})
	if err != nil {
		log.Error(fmt.Sprintf("Error in creating new qdrant client: %v", err))
		return nil, err
	}
	return client, nil
}
