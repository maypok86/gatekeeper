package storage

import (
	"context"
	"net"
)

type IP interface {
	Contains(ctx context.Context, ip net.IP) bool
	AddIP(ctx context.Context, ip net.IP)
	AddSubnet(ctx context.Context, subnet *net.IPNet)
	GetAll(ctx context.Context) []string
	RemoveIP(ctx context.Context, ip string)
	RemoveSubnet(ctx context.Context, subnet string)
}
