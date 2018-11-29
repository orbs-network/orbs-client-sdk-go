package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestExtractLatestTagFromDockerHubResponse(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
	}{
		{
			name:"HappyFlow",
			input: `{"count": 2, "next": null, "previous": null, "results": [{"name": "v0.4.2", "full_size": 126289039, "images": [{"size": 126289039, "architecture": "amd64", "variant": null, "features": null, "os": "linux", "os_version": null, "os_features": null}], "id": 40717129, "repository": 6341803, "creator": 4691149, "last_updater": 4691149, "last_updated": "2018-11-26T14:02:32.560968Z", "image_id": null, "v2": true}, {"name": "v0.7.0", "full_size": 126289039, "images": [{"size": 126289039, "architecture": "amd64", "variant": null, "features": null, "os": "linux", "os_version": null, "os_features": null}], "id": 40686462, "repository": 6341803, "creator": 4691149, "last_updater": 4691149, "last_updated": "2018-11-26T14:02:27.625449Z", "image_id": null, "v2": true}]}`,
			expected: "v0.7.0",
		},
		{
			name:"HappyFlowReversed",
			input: `{"count": 2, "next": null, "previous": null, "results": [{"name": "v1.2.3", "full_size": 126289039, "images": [{"size": 126289039, "architecture": "amd64", "variant": null, "features": null, "os": "linux", "os_version": null, "os_features": null}], "id": 40717129, "repository": 6341803, "creator": 4691149, "last_updater": 4691149, "last_updated": "2018-11-26T14:02:32.560968Z", "image_id": null, "v2": true}, {"name": "v0.7.0", "full_size": 126289039, "images": [{"size": 126289039, "architecture": "amd64", "variant": null, "features": null, "os": "linux", "os_version": null, "os_features": null}], "id": 40686462, "repository": 6341803, "creator": 4691149, "last_updater": 4691149, "last_updated": "2018-11-26T14:02:27.625449Z", "image_id": null, "v2": true}]}`,
			expected: "v1.2.3",
		},
		{
			name:"Empty",
			input: ``,
			expected: "",
		},
		{
			name:"NoResults",
			input: `{"count": 0, "next": null, "previous": null, "results": []}`,
			expected: "",
		},
		{
			name:"Corrupt",
			input: `{"count": 2, "next": null, "previous": null, "results": [{"name": "v0.4.2", "full_size": 126289039, "ima`,
			expected: "",
		},
		{
			name:"NonSemver",
			input: `{"count": 2, "next": null, "previous": null, "results": [{"name": "latest", "full_size": 126289039, "images": [{"size": 126289039, "architecture": "amd64", "variant": null, "features": null, "os": "linux", "os_version": null, "os_features": null}], "id": 40717129, "repository": 6341803, "creator": 4691149, "last_updater": 4691149, "last_updated": "2018-11-26T14:02:32.560968Z", "image_id": null, "v2": true}, {"name": "v0.7.0", "full_size": 126289039, "images": [{"size": 126289039, "architecture": "amd64", "variant": null, "features": null, "os": "linux", "os_version": null, "os_features": null}], "id": 40686462, "repository": 6341803, "creator": 4691149, "last_updater": 4691149, "last_updated": "2018-11-26T14:02:27.625449Z", "image_id": null, "v2": true}]}`,
			expected: "v0.7.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := extractLatestTagFromDockerHubResponse([]byte(tt.input))
			if tt.expected == "" {
				require.Error(t, err, "extract should return an error")
			} else {
				require.NoError(t, err, "extract should not return an error")
				require.Equal(t, tt.expected, tag, "extracted tag should match")
			}
		})
	}
}
