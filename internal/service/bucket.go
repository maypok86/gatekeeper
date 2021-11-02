package service

import (
	"context"

	"github.com/maypok86/gatekeeper/internal/storage"
)

type Bucket interface {
	Allow(name string) bool
	Remove(name string)
	HasBucket(name string) bool
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

func (bs *bucketService) Allow(name string) bool {
	return bs.bucket.Allow(bs.ctx, name)
}

func (bs *bucketService) Remove(name string) {
	bs.bucket.Remove(bs.ctx, name)
}

func (bs *bucketService) HasBucket(name string) bool {
	return bs.bucket.HasBucket(bs.ctx, name)
}
