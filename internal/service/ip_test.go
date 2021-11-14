package service

import (
	"context"
	"testing"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/stretchr/testify/require"
)

func TestNewAllowIPService(t *testing.T) {
	ctx := context.Background()
	service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
	require.NoError(t, err)
	require.NotNil(t, service)
}

func TestIPService_Contains_Found(t *testing.T) {
	tests := []struct {
		name   string
		ips    []string
		search string
	}{
		{
			name:   "single ip",
			ips:    []string{"192.168.1.1"},
			search: "192.168.1.1",
		},
		{
			name:   "single subnet",
			ips:    []string{"192.168.1.0/24"},
			search: "192.168.1.1",
		},
		{
			name:   "subnet and ip(2)",
			ips:    []string{"192.168.1.0/24", "172.168.10.10", "172.168.10.12"},
			search: "172.168.10.12",
		},
		{
			name:   "subnet(2) and ip",
			ips:    []string{"192.168.1.0/24", "172.168.10.0/16", "10.10.10.12"},
			search: "172.168.10.12",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
			require.NoError(t, err)
			for _, ip := range test.ips {
				err := service.Add(ctx, ip)
				require.NoError(t, err, "Unable add ip/subnet for test data")
			}
			found, err := service.Contains(ctx, test.search)
			require.NoError(t, err, "Error on check ip/subnet")
			require.True(t, found)
		})
	}
}

func TestIPService_Contains_NotFound(t *testing.T) {
	tests := []struct {
		name   string
		ips    []string
		search string
	}{
		{
			name:   "single ip",
			ips:    []string{"192.168.1.2"},
			search: "192.168.1.1",
		},
		{
			name:   "single subnet",
			ips:    []string{"192.168.1.0/24"},
			search: "192.168.2.1",
		},
		{
			name:   "subnet and ip(2)",
			ips:    []string{"192.168.1.0/24", "172.168.10.10", "172.168.10.12"},
			search: "10.10.10.12",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
			require.NoError(t, err)
			for _, ip := range test.ips {
				err := service.Add(ctx, ip)
				require.NoError(t, err, "Unable add ip/subnet for test data")
			}
			found, err := service.Contains(ctx, test.search)
			require.NoError(t, err, "Error on check ip/subnet")
			require.False(t, found)
		})
	}
}

func TestIPService_Contains_CheckInvalidIP(t *testing.T) {
	ctx := context.Background()
	service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
	require.NoError(t, err)
	_, err = service.Contains(ctx, "192.168.1")
	require.ErrorIs(t, err, ErrInvalidIpv4Address)
}

func TestIPService_Add_InvalidIP(t *testing.T) {
	ctx := context.Background()
	service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
	require.NoError(t, err)
	err = service.Add(ctx, "192.168.1.")
	require.ErrorIs(t, err, ErrInvalidIpv4Address)
}

func TestIPService_Add_InvalidSubnet(t *testing.T) {
	ctx := context.Background()
	service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
	require.NoError(t, err)
	err = service.Add(ctx, "192.168.1./50")
	require.ErrorIs(t, err, ErrInvalidIpv4Subnet)
}

func TestIPService_Remove_Exists(t *testing.T) {
	tests := []struct {
		name     string
		tableIP  string
		searchIP string
	}{
		{
			name:     "remove subnet",
			tableIP:  "10.10.1.0/24",
			searchIP: "10.10.1.10",
		},
		{
			name:     "remove ip",
			tableIP:  "10.10.1.10",
			searchIP: "10.10.1.10",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
			require.NoError(t, err)
			err = service.Add(ctx, test.tableIP)
			require.NoError(t, err)

			f, err := service.Contains(ctx, test.searchIP)
			require.NoError(t, err)
			require.True(t, f)

			service.Remove(ctx, test.tableIP)

			f, err = service.Contains(ctx, test.searchIP)
			require.NoError(t, err)
			require.False(t, f)
		})
	}
}

func TestIPService_Remove_NotExists(t *testing.T) {
	tests := []struct {
		name    string
		tableIP string
	}{
		{
			name:    "remove subnet",
			tableIP: "10.10.1.0/24",
		},
		{
			name:    "remove ip",
			tableIP: "10.10.1.10",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			service, err := NewIP(ctx, &config.Config{DBType: "inmemory"})
			require.NoError(t, err)
			require.NotPanics(t, func() {
				service.Remove(ctx, test.tableIP)
			})
		})
	}
}
