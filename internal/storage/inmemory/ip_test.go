package inmemory

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAllowIPStorage(t *testing.T) {
	st := NewIPStorage()
	require.NotNil(t, st)
}

func TestIPStorage_Contains_Found(t *testing.T) {
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
	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := NewIPStorage()
			for _, ip := range test.ips {
				err := st.Add(ctx, ip)
				require.NoError(t, err, "Unable add ip/subnet for test data")
			}
			found, err := st.Contains(ctx, test.search)
			require.NoError(t, err, "Error on check ip/subnet")
			require.True(t, found)
		})
	}
}

func TestIPStorage_Contains_NotFound(t *testing.T) {
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
	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := NewIPStorage()
			for _, ip := range test.ips {
				err := st.Add(ctx, ip)
				require.NoError(t, err, "Unable add ip/subnet for test data")
			}
			found, err := st.Contains(ctx, test.search)
			require.NoError(t, err, "Error on check ip/subnet")
			require.False(t, found)
		})
	}
}

func TestIPStorage_Contains_CheckInvalidIP(t *testing.T) {
	st := NewIPStorage()
	ctx := context.Background()
	_, err := st.Contains(ctx, "192.168.1")
	require.ErrorIs(t, err, ErrInvalidIpv4Address)
}

func TestIPStorage_Add_InvalidIP(t *testing.T) {
	st := NewIPStorage()
	ctx := context.Background()
	err := st.Add(ctx, "192.168.1.")
	require.ErrorIs(t, err, ErrInvalidIpv4Address)
}

func TestIPStorage_Add_InvalidSubnet(t *testing.T) {
	st := NewIPStorage()
	ctx := context.Background()
	err := st.Add(ctx, "192.168.1./50")
	require.ErrorIs(t, err, ErrInvalidIpv4Subnet)
}

func TestIPStorage_GetAll_Empty(t *testing.T) {
	st := NewIPStorage()
	ctx := context.Background()
	require.Len(t, st.GetAll(ctx), 0)
}

func TestIPStorage_GetAll_NotEmpty(t *testing.T) {
	st := NewIPStorage()
	ctx := context.Background()
	ips := []string{"192.168.1.12", "10.10.0.0/24"}
	sort.Strings(ips)
	for _, ip := range ips {
		err := st.Add(ctx, ip)
		require.NoError(t, err)
	}
	require.Equal(t, ips, st.GetAll(ctx))
}

func TestIPStorage_Remove_Exists(t *testing.T) {
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
	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := NewIPStorage()
			err := st.Add(ctx, test.tableIP)
			require.NoError(t, err)

			f, err := st.Contains(ctx, test.searchIP)
			require.NoError(t, err)
			require.True(t, f)

			st.Remove(ctx, test.tableIP)

			f, err = st.Contains(ctx, test.searchIP)
			require.NoError(t, err)
			require.False(t, f)
		})
	}
}

func TestIPStorage_Remove_NotExists(t *testing.T) {
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
	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := NewIPStorage()
			require.NotPanics(t, func() {
				st.Remove(ctx, test.tableIP)
			})
		})
	}
}
