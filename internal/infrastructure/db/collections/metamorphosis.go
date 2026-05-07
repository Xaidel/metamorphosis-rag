package collections

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

const collectionName = "metamorphosis"

func NewCollection(ctx context.Context, client *qdrant.Client) error {
	exists, err := client.CollectionExists(ctx, collectionName)
	if err != nil {
		return fmt.Errorf("Error in searching  collection %v", err)
	}

	if exists {
		err := client.CreateCollection(ctx, &qdrant.CreateCollection{
			CollectionName: collectionName,
			VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
				Size:     768,
				Distance: qdrant.Distance_Cosine,
			}),
		})

		if err != nil {
			return err
		}
	}
	return nil
}
