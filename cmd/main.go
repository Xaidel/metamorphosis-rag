package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/xaidel/metamorphosis-rag/internal/infrastructure/di"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error here")
		panic(err)
	}
	fmt.Println("Application exited successfully")
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	application, err := di.Bootstrap(ctx)
	if err != nil {
		return fmt.Errorf("bootstrapping application: %w", err)
	}

	defer func() {
		_ = application.Shutdown()
	}()

	return nil
}
