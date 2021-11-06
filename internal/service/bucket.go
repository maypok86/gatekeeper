package service

import (
	"context"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/internal/storage"
	"github.com/maypok86/gatekeeper/internal/storage/inmemory"
	"github.com/pkg/errors"
)

type Bucket interface {
	Allow(ctx context.Context, name string) bool
	Remove(ctx context.Context, name string)
	HasBucket(ctx context.Context, name string) bool
}

type bucketService struct {
	ctx    context.Context
	bucket storage.Bucket
}

func NewBucket(ctx context.Context, cfg *config.Config) (Bucket, error) {
	var bs *bucketService
	switch cfg.DBType {
	case "inmemory":
		bs = &bucketService{
			ctx:    ctx,
			bucket: inmemory.NewBucketStorage(cfg.Bucket),
		}
	default:
		return nil, errors.New("unknown type db")
	}
	return bs, nil
}

func (bs *bucketService) Allow(ctx context.Context, name string) bool {
	return bs.bucket.Allow(ctx, name)
}

func (bs *bucketService) Remove(ctx context.Context, name string) {
	bs.bucket.Remove(ctx, name)
}

func (bs *bucketService) HasBucket(ctx context.Context, name string) bool {
	return bs.bucket.HasBucket(ctx, name)
}
