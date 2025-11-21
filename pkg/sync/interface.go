package sync

import (
	"context"
	"time"
)

// Syncer orchestrates data synchronization from multiple transport providers.
// It provides methods to sync all data or data from specific providers.
type Syncer interface {
	// SyncAll synchronizes data from all configured providers.
	// It continues processing even if one provider fails, logging errors.
	SyncAll(ctx context.Context) error

	// SyncProvider synchronizes data from a specific provider.
	SyncProvider(ctx context.Context, provider Provider) error

	// StartPeriodicSync runs synchronization on a schedule.
	// DEPRECATED: Use host cron instead for production (see docs/DEPLOYMENT.md).
	// This method is kept for backward compatibility and local development only.
	// It performs an initial sync immediately, then repeats every interval.
	// Blocks until context is cancelled.
	StartPeriodicSync(ctx context.Context, interval time.Duration)
}

// ProviderClient represents a generic external API client.
// This interface can be implemented by all provider clients for uniform handling.
type ProviderClient interface {
	// GetName returns the provider identifier (e.g., "gars", "aviasales", "rzd").
	GetName() string

	// HealthCheck verifies API connectivity and authentication.
	HealthCheck(ctx context.Context) error
}
