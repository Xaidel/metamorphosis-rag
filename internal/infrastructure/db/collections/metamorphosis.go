package collections

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/qdrant/go-client/qdrant"
	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/config"
)

func NewCollection(ctx context.Context, client *qdrant.Client, collection config.Collection) error {
	exists, err := client.CollectionExists(ctx, collection.Name)
	if err != nil {
		return fmt.Errorf("Error in searching  collection %v", err)
	}

	if !exists {
		err := client.CreateCollection(ctx, &qdrant.CreateCollection{
			CollectionName: collection.Name,
			VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
				Size:     768,
				Distance: qdrant.Distance_Cosine,
			}),
		})

		if err != nil {
			return err
		}
		log.Info(fmt.Sprintf("Collection %s successfully created", collection.Name))
	}
	log.Info(fmt.Sprintf("Collection %s is ready", collection.Name))
	return nil
}
