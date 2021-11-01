package inmemory

import (
	"context"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"

	"github.com/maypok86/gatekeeper/pkg/logger"
	"github.com/pkg/errors"
)

var (
	ErrInvalidIpv4Address = errors.New("invalid ip address")
	ErrInvalidIpv4Subnet  = errors.New("invalid ip subnet")
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

func (is *IPStorage) Contains(ctx context.Context, clientIP string) (bool, error) {
	select {
	case <-ctx.Done():
		logger.Warn("client ip contains cancel: ", clientIP)
		return false, nil
	default:
		ip := net.ParseIP(clientIP)
		if ip.To4() == nil {
			return false, fmt.Errorf("%w:%s is not valid ipv4 address", ErrInvalidIpv4Address, clientIP)
		}
		is.mutex.Lock()
		defer is.mutex.Unlock()

		if _, ok := is.ipSet[clientIP]; ok {
			return true, nil
		}
		for _, subnet := range is.subnetMap {
			if subnet.Contains(ip) {
				return true, nil
			}
		}
		return false, nil
	}
}

func (is *IPStorage) Add(ctx context.Context, ipOrSubnet string) error {
	select {
	case <-ctx.Done():
		logger.Warn("Add ip or subnet cancel: ", ipOrSubnet)
		return nil
	default:
		if isSubnet(ipOrSubnet) {
			return is.addSubnet(ipOrSubnet)
		}
		return is.addIP(ipOrSubnet)
	}
}

func isSubnet(ipOrSubnet string) bool {
	return strings.Contains(ipOrSubnet, "/")
}

func (is *IPStorage) addIP(clientIP string) error {
	ip := net.ParseIP(clientIP)
	if ip.To4() == nil {
		return fmt.Errorf("%w:%s is not valid ipv4 address", ErrInvalidIpv4Address, clientIP)
	}

	is.mutex.Lock()
	defer is.mutex.Unlock()
	is.ipSet[clientIP] = struct{}{}
	return nil
}

func (is *IPStorage) addSubnet(subnet string) error {
	_, cidr, err := net.ParseCIDR(subnet)
	if err != nil {
		return fmt.Errorf("%w: is not valid ipv4 subnet %s,  %s", ErrInvalidIpv4Subnet, subnet, err)
	}

	is.mutex.Lock()
	defer is.mutex.Unlock()
	is.subnetMap[subnet] = cidr
	return nil
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

func (is *IPStorage) Remove(ctx context.Context, ipOrSubnet string) {
	select {
	case <-ctx.Done():
		logger.Warn("Remove ip or subnet cancel: ", ipOrSubnet)
	default:
		is.mutex.Lock()
		defer is.mutex.Unlock()
		if isSubnet(ipOrSubnet) {
			delete(is.subnetMap, ipOrSubnet)
		} else {
			delete(is.ipSet, ipOrSubnet)
		}
	}
}
