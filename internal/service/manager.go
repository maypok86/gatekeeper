package service

import (
	"context"

	"github.com/maypok86/gatekeeper/internal/storage"
	"github.com/pkg/errors"
)

type Manager struct {
	Bucket Bucket
	IP     IP
}

func NewManager(ctx context.Context, st *storage.Storage) (*Manager, error) {
	if st == nil {
		return nil, errors.New("No storage provided")
	}
	return &Manager{
		Bucket: NewBucket(ctx, st.Bucket),
		IP:     NewIP(ctx, st.IP),
	}, nil
}
