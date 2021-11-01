package storage

import "context"

type IP interface {
	Contains(ctx context.Context, ip string) (bool, error)
	Add(ctx context.Context, ipOrSubnet string) error
	GetAll(ctx context.Context) []string
	Remove(ctx context.Context, ipOrSubnet string)
}
