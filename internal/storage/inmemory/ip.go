package inmemory

import (
	"context"
	"net"
	"sort"
	"sync"

	"github.com/maypok86/gatekeeper/pkg/logger"
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
		logger.Warn("client ip contains cancel: ", ip.String())
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
		logger.Warn("Add ip cancel: ", ip.String())
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		is.ipSet[ip.String()] = struct{}{}
	}
}

func (is *IPStorage) AddSubnet(ctx context.Context, subnet *net.IPNet) {
	select {
	case <-ctx.Done():
		logger.Warn("Add subnet cancel: ", subnet.String())
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		is.subnetMap[subnet.String()] = subnet
	}
}

func (is *IPStorage) GetAll(ctx context.Context) []string {
	select {
	case <-ctx.Done():
		logger.Warn("Get all ips cancel")
		return []string{}
	default:
		ips := make([]string, 0)
		for ip := range is.ipSet {
			ips = append(ips, ip)
		}
		for subnet := range is.subnetMap {
			ips = append(ips, subnet)
		}
		sort.Strings(ips)
		return ips
	}
}

func (is *IPStorage) RemoveIP(ctx context.Context, ip string) {
	select {
	case <-ctx.Done():
		logger.Warn("Remove ip cancel: ", ip)
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		delete(is.ipSet, ip)
	}
}

func (is *IPStorage) RemoveSubnet(ctx context.Context, subnet string) {
	select {
	case <-ctx.Done():
		logger.Warn("Remove subnet cancel: ", subnet)
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		delete(is.subnetMap, subnet)
	}
}
