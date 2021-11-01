package storage

import (
	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/internal/storage/inmemory"
	"github.com/pkg/errors"
)

type Storage struct {
	Bucket Bucket
	IP     IP
}

func New(cfg *config.Bucket) (*Storage, error) {
	var s *Storage
	switch cfg.Type {
	case "inmemory":
		s = &Storage{
			Bucket: inmemory.NewBucketStorage(cfg),
			IP:     inmemory.NewIPStorage(),
		}
	default:
		return nil, errors.New("unknown type db")
	}

	return s, nil
}
