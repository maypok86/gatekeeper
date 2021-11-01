package storage

import "context"

type Bucket interface {
	Allow(ctx context.Context, name string) bool
	Remove(ctx context.Context, name string)
	HasBucket(ctx context.Context, name string) bool
}
