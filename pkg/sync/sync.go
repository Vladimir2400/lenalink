package sync

import (
	"context"
	"time"

	"github.com/lenalink/backend/internal/repository"
	"github.com/lenalink/backend/pkg/sync/api/aviasales"
	"github.com/lenalink/backend/pkg/sync/api/gars"
	"github.com/lenalink/backend/pkg/sync/api/rzd"
)

// Package sync provides data synchronization from external transport providers
// (GARS, Aviasales, RZD) to the LenaLink database.

// New creates a new Syncer with all required dependencies.
func New(
	garsClient *gars.Client,
	aviasalesClient *aviasales.Client,
	rzdClient *rzd.MockClient,
	stopRepo repository.StopRepository,
	segmentRepo repository.SegmentRepository,
) Syncer {
	return &service{
		garsClient:      garsClient,
		aviasalesClient: aviasalesClient,
		rzdClient:       rzdClient,
		stopRepo:        stopRepo,
		segmentRepo:     segmentRepo,
	}
}

// Provider represents an external transport data provider.
type Provider string

const (
	// ProviderGARS represents АвиБус (bus routes) provider.
	ProviderGARS Provider = "gars"
	// ProviderAviasales represents flight pricing provider.
	ProviderAviasales Provider = "aviasales"
	// ProviderRZD represents Russian Railways provider.
	ProviderRZD Provider = "rzd"
)

// SyncOptions configures synchronization behavior.
type SyncOptions struct {
	// PeriodicInterval defines how often to run sync (0 = run once).
	PeriodicInterval time.Duration

	// Providers to sync (nil = all providers).
	Providers []Provider

	// CleanupOlderThan removes segments older than this duration.
	CleanupOlderThan time.Duration
}

// DefaultOptions returns recommended sync configuration.
func DefaultOptions() SyncOptions {
	return SyncOptions{
		PeriodicInterval: 6 * time.Hour,
		Providers:        nil, // sync all
		CleanupOlderThan: 7 * 24 * time.Hour,
	}
}

// RunSync is a helper function that creates a Syncer and runs synchronization once.
func RunSync(ctx context.Context,
	garsClient *gars.Client,
	aviasalesClient *aviasales.Client,
	rzdClient *rzd.MockClient,
	stopRepo repository.StopRepository,
	segmentRepo repository.SegmentRepository,
) error {
	syncer := New(garsClient, aviasalesClient, rzdClient, stopRepo, segmentRepo)
	return syncer.SyncAll(ctx)
}
