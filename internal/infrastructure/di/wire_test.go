package di_test

import (
	"testing"

	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/di"
)

func TestApplicationShutdownNil(t *testing.T) {
	t.Parallel()
	var app *di.Application
	if err := app.Shutdown(); err != nil {
		t.Fatalf("There's an error when shutting down the app %v", err)
	}
}
