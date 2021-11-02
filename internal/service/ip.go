package service

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/maypok86/gatekeeper/internal/storage"
	"github.com/pkg/errors"
)

var (
	ErrInvalidIpv4Address = errors.New("invalid ip address")
	ErrInvalidIpv4Subnet  = errors.New("invalid ip subnet")
)

type IP interface {
	Contains(clientIP string) (bool, error)
	Add(ipOrSubnet string) error
	GetAll() []string
	Remove(ipOrSubnet string)
}

type ipService struct {
	ctx context.Context
	ip  storage.IP
}

func NewIP(ctx context.Context, ip storage.IP) IP {
	return &ipService{
		ctx: ctx,
		ip:  ip,
	}
}

func (is *ipService) Contains(clientIP string) (bool, error) {
	ip := net.ParseIP(clientIP)
	if ip.To4() == nil {
		return false, fmt.Errorf("%w:%s is not valid ipv4 address", ErrInvalidIpv4Address, clientIP)
	}
	return is.ip.Contains(is.ctx, ip), nil
}

func isSubnet(ipOrSubnet string) bool {
	return strings.Contains(ipOrSubnet, "/")
}

func (is *ipService) Add(ipOrSubnet string) error {
	if isSubnet(ipOrSubnet) {
		_, cidr, err := net.ParseCIDR(ipOrSubnet)
		if err != nil {
			return fmt.Errorf("%w: is not valid ipv4 subnet %s,  %s", ErrInvalidIpv4Subnet, ipOrSubnet, err)
		}
		is.ip.AddSubnet(is.ctx, cidr)
	} else {
		ip := net.ParseIP(ipOrSubnet)
		if ip.To4() == nil {
			return fmt.Errorf("%w:%s is not valid ipv4 address", ErrInvalidIpv4Address, ipOrSubnet)
		}
		is.ip.AddIP(is.ctx, ip)
	}
	return nil
}

func (is *ipService) GetAll() []string {
	return is.ip.GetAll(is.ctx)
}

func (is *ipService) Remove(ipOrSubnet string) {
	if isSubnet(ipOrSubnet) {
		is.ip.RemoveSubnet(is.ctx, ipOrSubnet)
	} else {
		is.ip.RemoveIP(is.ctx, ipOrSubnet)
	}
}
