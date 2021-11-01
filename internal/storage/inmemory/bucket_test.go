package inmemory

import (
	"context"
	"testing"
	"time"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/stretchr/testify/require"
)

const (
	testBucketName = "test"
	defaultTTL     = 60
)

var ctx = context.Background()

func changeConfigRate(rate int) {
	config.Get().Bucket.Rate = rate
}

func TestNewBucketStorage(t *testing.T) {
	rate := 200
	changeConfigRate(rate)
	b := NewBucketStorage()
	require.NotNil(t, b)
	require.Equal(t, rate, b.rate)
}

func TestBucketStorage_Allow(t *testing.T) {
	changeConfigRate(2)
	st := NewBucketStorage()
	require.True(t, st.Allow(ctx, "test"))
	require.True(t, st.Allow(ctx, "test"))
	require.False(t, st.Allow(ctx, "test"))
	require.True(t, st.Allow(ctx, "test2"))
}

func TestBucketStorage_Clean(t *testing.T) {
	changeConfigRate(2)
	st := NewBucketStorage()
	st.buckets["f1"] = &bucket{
		lastAccess: time.Now().Unix() - defaultTTL - 10,
	}
	st.buckets["f2"] = &bucket{
		lastAccess: time.Now().Unix() - defaultTTL + 10,
	}

	st.clean()
	require.Len(t, st.buckets, 1)
	_, ok := st.buckets["f2"]
	require.True(t, ok)
}

func TestBucketStorage_Remove(t *testing.T) {
	changeConfigRate(2)
	st := NewBucketStorage()
	require.True(t, st.Allow(ctx, testBucketName))
	require.True(t, st.Allow(ctx, testBucketName))
	require.False(t, st.Allow(ctx, testBucketName))
	st.Remove(ctx, testBucketName)
	require.True(t, st.Allow(ctx, testBucketName))
	require.True(t, st.Allow(ctx, testBucketName))
	require.False(t, st.Allow(ctx, testBucketName))
}

func TestBucketStorage_CleanByTimer(t *testing.T) {
	config.Get().Bucket = &config.Bucket{
		Rate:          2,
		TTL:           1,
		CleanInterval: 1010 * time.Millisecond,
	}

	st := NewBucketStorage()
	require.True(t, st.Allow(ctx, testBucketName))
	require.True(t, st.HasBucket(ctx, testBucketName))
	cond := func() bool {
		return st.HasBucket(ctx, testBucketName)
	}
	require.Eventually(t, cond, 5*time.Second, 100*time.Millisecond)
}

func TestBucketStorage_TokenExpire(t *testing.T) {
	config.Get().Bucket = &config.Bucket{
		Rate: 2,
		TTL:  1,
	}

	st := NewBucketStorage()
	require.True(t, st.Allow(ctx, testBucketName))
	require.True(t, st.Allow(ctx, testBucketName))
	time.Sleep(config.Get().Bucket.TTL * time.Second)
	require.True(t, st.Allow(ctx, testBucketName))
}

func TestBucketStorage_HasBucket_Exists(t *testing.T) {
	changeConfigRate(2)
	st := NewBucketStorage()
	require.True(t, st.Allow(ctx, testBucketName))
	require.True(t, st.HasBucket(ctx, testBucketName))
}

func TestBucketStorage_HasBucket_NotExists(t *testing.T) {
	changeConfigRate(2)
	reqLimiter := NewBucketStorage()
	require.False(t, reqLimiter.HasBucket(ctx, testBucketName))
}
