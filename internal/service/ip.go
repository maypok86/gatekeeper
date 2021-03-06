package service

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/internal/storage"
	"github.com/maypok86/gatekeeper/internal/storage/inmemory"
	"github.com/pkg/errors"
)

var (
	ErrInvalidIpv4Address = errors.New("invalid ip address")
	ErrInvalidIpv4Subnet  = errors.New("invalid ip subnet")
)

type IP interface {
	Contains(ctx context.Context, clientIP string) (bool, error)
	Add(ctx context.Context, ipOrSubnet string) error
	Remove(ctx context.Context, ipOrSubnet string)
}

type ipService struct {
	ctx context.Context
	ip  storage.IP
}

func NewIP(ctx context.Context, cfg *config.Config) (IP, error) {
	var is *ipService
	switch cfg.DBType {
	case "inmemory":
		is = &ipService{
			ctx: ctx,
			ip:  inmemory.NewIPStorage(),
		}
	default:
		return nil, errors.New("unknown type db")
	}
	return is, nil
}

func (is *ipService) Contains(ctx context.Context, clientIP string) (bool, error) {
	ip := net.ParseIP(clientIP)
	if ip.To4() == nil {
		return false, fmt.Errorf("%w:%s is not valid ipv4 address", ErrInvalidIpv4Address, clientIP)
	}
	return is.ip.Contains(ctx, ip), nil
}

func isSubnet(ipOrSubnet string) bool {
	return strings.Contains(ipOrSubnet, "/")
}

func (is *ipService) Add(ctx context.Context, ipOrSubnet string) error {
	if isSubnet(ipOrSubnet) {
		_, cidr, err := net.ParseCIDR(ipOrSubnet)
		if err != nil {
			return fmt.Errorf("%w: is not valid ipv4 subnet %s,  %s", ErrInvalidIpv4Subnet, ipOrSubnet, err)
		}
		is.ip.AddSubnet(ctx, cidr)
	} else {
		ip := net.ParseIP(ipOrSubnet)
		if ip.To4() == nil {
			return fmt.Errorf("%w:%s is not valid ipv4 address", ErrInvalidIpv4Address, ipOrSubnet)
		}
		is.ip.AddIP(ctx, ip)
	}
	return nil
}

func (is *ipService) Remove(ctx context.Context, ipOrSubnet string) {
	if isSubnet(ipOrSubnet) {
		is.ip.RemoveSubnet(ctx, ipOrSubnet)
	} else {
		is.ip.RemoveIP(ctx, ipOrSubnet)
	}
}
