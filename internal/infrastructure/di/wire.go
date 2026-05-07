package di

import (
	"context"

	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/config"
	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/db"
	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/db/collections"
)

type Application struct {
	Config *config.AppConfig
}

func Bootstrap(ctx context.Context) (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	vector_storage, err := db.NewVectorStorage(cfg.Storage)
	if err != nil {
		vector_storage.Close()
		return nil, err
	}

	err = collections.NewCollection(ctx, vector_storage, cfg.Collection)
	if err != nil {
		return nil, err
	}

	return &Application{
		Config: cfg,
	}, nil
}

func (a *Application) Shutdown() error {
	if a == nil {
		return nil
	}
	//TODO: Add any necessary cleanup logic here (e.g., closing database connections, stopping background tasks, etc.)
	var cleanupErr error

	return cleanupErr
}
