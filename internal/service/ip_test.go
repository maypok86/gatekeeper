package service

import (
	"context"
	"sort"
	"testing"

	"github.com/maypok86/gatekeeper/internal/storage/inmemory"
	"github.com/stretchr/testify/require"
)

func TestNewAllowIPService(t *testing.T) {
	service := NewIP(context.Background(), inmemory.NewIPStorage())
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
			service := NewIP(context.Background(), inmemory.NewIPStorage())
			for _, ip := range test.ips {
				err := service.Add(ip)
				require.NoError(t, err, "Unable add ip/subnet for test data")
			}
			found, err := service.Contains(test.search)
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
			service := NewIP(context.Background(), inmemory.NewIPStorage())
			for _, ip := range test.ips {
				err := service.Add(ip)
				require.NoError(t, err, "Unable add ip/subnet for test data")
			}
			found, err := service.Contains(test.search)
			require.NoError(t, err, "Error on check ip/subnet")
			require.False(t, found)
		})
	}
}

func TestIPService_Contains_CheckInvalidIP(t *testing.T) {
	service := NewIP(context.Background(), inmemory.NewIPStorage())
	_, err := service.Contains("192.168.1")
	require.ErrorIs(t, err, ErrInvalidIpv4Address)
}

func TestIPService_Add_InvalidIP(t *testing.T) {
	service := NewIP(context.Background(), inmemory.NewIPStorage())
	err := service.Add("192.168.1.")
	require.ErrorIs(t, err, ErrInvalidIpv4Address)
}

func TestIPService_Add_InvalidSubnet(t *testing.T) {
	service := NewIP(context.Background(), inmemory.NewIPStorage())
	err := service.Add("192.168.1./50")
	require.ErrorIs(t, err, ErrInvalidIpv4Subnet)
}

func TestIPService_GetAll_Empty(t *testing.T) {
	service := NewIP(context.Background(), inmemory.NewIPStorage())
	require.Len(t, service.GetAll(), 0)
}

func TestIPService_GetAll_NotEmpty(t *testing.T) {
	service := NewIP(context.Background(), inmemory.NewIPStorage())
	ips := []string{"192.168.1.12", "10.10.0.0/24"}
	sort.Strings(ips)
	for _, ip := range ips {
		err := service.Add(ip)
		require.NoError(t, err)
	}
	require.Equal(t, ips, service.GetAll())
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
			service := NewIP(context.Background(), inmemory.NewIPStorage())
			err := service.Add(test.tableIP)
			require.NoError(t, err)

			f, err := service.Contains(test.searchIP)
			require.NoError(t, err)
			require.True(t, f)

			service.Remove(test.tableIP)

			f, err = service.Contains(test.searchIP)
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
			service := NewIP(context.Background(), inmemory.NewIPStorage())
			require.NotPanics(t, func() {
				service.Remove(test.tableIP)
			})
		})
	}
}