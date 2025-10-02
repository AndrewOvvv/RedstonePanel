package minedb

import "context"

// MineDB is a minimal key-value/document database abstraction.
// Implementations are concurrency-safe.
type MineDB interface {
	// Get fetches a value by bucket and key. Returns (nil, nil) if not found.
	Get(ctx context.Context, bucket, key string) (any, error)

	// Set stores a value by bucket and key, replacing any existing value.
	Set(ctx context.Context, bucket, key string, value any) error

	// Delete deletes a value by bucket and key. Missing keys are ignored.
	Delete(ctx context.Context, bucket, key string) error

	// Scan iterates keys in a bucket.
	Scan(ctx context.Context, bucket string, fn func(k string, v any) (cont bool, err error)) error
}
