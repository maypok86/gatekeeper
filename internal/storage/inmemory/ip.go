package inmemory

import (
	"context"
	"net"
	"sync"

	"go.uber.org/zap"
)

type IPStorage struct {
	subnetMap map[string]*net.IPNet
	ipSet     map[string]struct{}
	mutex     sync.Mutex
}

func NewIPStorage() *IPStorage {
	return &IPStorage{
		subnetMap: make(map[string]*net.IPNet),
		ipSet:     make(map[string]struct{}),
	}
}

func (is *IPStorage) Contains(ctx context.Context, ip net.IP) bool {
	select {
	case <-ctx.Done():
		zap.L().Warn("client ip contains cancel: ", zap.String("ip", ip.String()))
		return false
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()

		if _, ok := is.ipSet[ip.String()]; ok {
			return true
		}
		for _, subnet := range is.subnetMap {
			if subnet.Contains(ip) {
				return true
			}
		}
		return false
	}
}

func (is *IPStorage) AddIP(ctx context.Context, ip net.IP) {
	select {
	case <-ctx.Done():
		zap.L().Warn("Add ip cancel: ", zap.String("ip", ip.String()))
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		is.ipSet[ip.String()] = struct{}{}
	}
}

func (is *IPStorage) AddSubnet(ctx context.Context, subnet *net.IPNet) {
	select {
	case <-ctx.Done():
		zap.L().Warn("Add subnet cancel: ", zap.String("subnet", subnet.String()))
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		is.subnetMap[subnet.String()] = subnet
	}
}

func (is *IPStorage) RemoveIP(ctx context.Context, ip string) {
	select {
	case <-ctx.Done():
		zap.L().Warn("Remove ip cancel: ", zap.String("ip", ip))
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		delete(is.ipSet, ip)
	}
}

func (is *IPStorage) RemoveSubnet(ctx context.Context, subnet string) {
	select {
	case <-ctx.Done():
		zap.L().Warn("Remove subnet cancel: ", zap.String("subnet", subnet))
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		delete(is.subnetMap, subnet)
	}
}
