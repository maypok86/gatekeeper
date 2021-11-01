package inmemory

import (
	"context"
	"sync"
	"time"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/pkg/logger"
	"golang.org/x/time/rate"
)

type BucketStorage struct {
	mutex   sync.Mutex
	ttl     time.Duration
	rate    int
	buckets map[string]*bucket
}

type bucket struct {
	limiter    *rate.Limiter
	lastAccess int64
	mutex      sync.Mutex
}

func NewBucketStorage(cfg *config.Bucket) *BucketStorage {
	st := &BucketStorage{
		ttl:     cfg.TTL,
		rate:    cfg.Rate,
		buckets: make(map[string]*bucket),
	}

	go func() {
		for range time.Tick(cfg.CleanInterval * time.Second) {
			st.clean()
		}
	}()

	return st
}

func (ib *BucketStorage) Allow(ctx context.Context, name string) bool {
	select {
	case <-ctx.Done():
		logger.Warn("bucket allow cancel: ", name)
		return false
	default:
		return ib.getBucket(name).Allow()
	}
}

func (ib *BucketStorage) getBucket(name string) *bucket {
	ib.mutex.Lock()
	defer ib.mutex.Unlock()

	b, ok := ib.buckets[name]
	if !ok {
		limiter := rate.NewLimiter(rate.Every(ib.ttl*time.Second), ib.rate)
		b = &bucket{
			limiter: limiter,
		}
		ib.buckets[name] = b
	}

	return b
}

func (ib *BucketStorage) Remove(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		logger.Warn("bucket remove cancel: ", name)
	default:
		ib.mutex.Lock()
		defer ib.mutex.Unlock()
		delete(ib.buckets, name)
	}
}

func (b *bucket) Allow() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.lastAccess = time.Now().Unix()
	return b.limiter.Allow()
}

func (ib *BucketStorage) clean() {
	ib.mutex.Lock()
	defer ib.mutex.Unlock()
	expire := time.Now().Unix() - int64(ib.ttl)
	for i, v := range ib.buckets {
		if expire > v.lastAccess {
			delete(ib.buckets, i)
		}
	}
}

func (ib *BucketStorage) HasBucket(ctx context.Context, name string) bool {
	select {
	case <-ctx.Done():
		logger.Warn("bucket has name cancel: ", name)
		return false
	default:
		ib.mutex.Lock()
		defer ib.mutex.Unlock()
		_, ok := ib.buckets[name]
		return ok
	}
}
