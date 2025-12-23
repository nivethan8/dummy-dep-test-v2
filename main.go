package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// generateID generates a simple random ID
func generateID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	defer logger.Sync()

	logger.Info("Starting dummy application",
		zap.String("app_id", generateID()),
		zap.Time("started_at", time.Now()),
	)

	ctx := context.Background()
	if err := run(ctx, logger); err != nil {
		logger.Fatal("Application failed", zap.Error(err))
	}

	logger.Info("Application completed successfully")
}

func run(ctx context.Context, logger *zap.Logger) error {
	g, ctx := errgroup.WithContext(ctx)

	// Simulate some concurrent work
	for i := 0; i < 3; i++ {
		taskID := i
		g.Go(func() error {
			return processTask(ctx, logger, taskID)
		})
	}

	return g.Wait()
}

func processTask(ctx context.Context, logger *zap.Logger, taskID int) error {
	taskUUID := generateID()
	logger.Info("Processing task",
		zap.Int("task_id", taskID),
		zap.String("task_uuid", taskUUID),
	)

	// Simulate some work
	time.Sleep(100 * time.Millisecond)

	logger.Info("Task completed",
		zap.Int("task_id", taskID),
		zap.String("task_uuid", taskUUID),
	)

	return nil
}
