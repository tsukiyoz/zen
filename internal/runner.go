package internal

import (
	"context"
	"log/slog"
	"time"
)

func Run(ctx context.Context) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Info("runner stoping...")

			// simulate clean logic
			time.Sleep(3 * time.Second)
			return nil
		case <-ticker.C:
			slog.Info("runner working...")
		}
	}
}
