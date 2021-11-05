package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	expects := "gk version 1.4.0 (2020-12-15)\nhttps://github.com/maypok86/gatekeeper/releases/tag/v1.4.0\n"
	got := format("1.4.0", "2020-12-15")
	require.Equal(t, expects, got)
}

func TestChangelogURL(t *testing.T) {
	tag := "0.3.2"
	url := "https://github.com/maypok86/gatekeeper/releases/tag/v0.3.2"
	result := changelogURL(tag)
	require.Equal(t, url, result)

	tag = "v0.3.2"
	url = "https://github.com/maypok86/gatekeeper/releases/tag/v0.3.2"
	result = changelogURL(tag)
	require.Equal(t, url, result)

	tag = "0.3.2-pre.1"
	url = "https://github.com/maypok86/gatekeeper/releases/tag/v0.3.2-pre.1"
	result = changelogURL(tag)
	require.Equal(t, url, result)

	tag = "0.3.5-90-gdd3f0e0"
	url = "https://github.com/maypok86/gatekeeper/releases/latest"
	result = changelogURL(tag)
	require.Equal(t, url, result)

	tag = "deadbeef"
	url = "https://github.com/maypok86/gatekeeper/releases/latest"
	result = changelogURL(tag)
	require.Equal(t, url, result)
}
