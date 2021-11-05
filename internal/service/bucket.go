package service

import (
	"context"

	"github.com/maypok86/gatekeeper/internal/storage"
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

func NewBucket(ctx context.Context, bucket storage.Bucket) Bucket {
	return &bucketService{
		ctx:    ctx,
		bucket: bucket,
	}
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
