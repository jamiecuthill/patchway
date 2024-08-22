package patchway_test

import (
	"testing"

	"github.com/jamiecuthill/patchway"
	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		path     []string
		expected string
	}{
		{
			name:     "General case",
			value:    "new value",
			path:     []string{"a", "b", "c"},
			expected: `{"op": "replace", "path": "/a/b/c", "value": "new value"}`,
		},
		{
			// See https://jsonpatch.com/#json-pointer
			name:     "Special characters",
			value:    "new value",
			path:     []string{"foo/bar~"},
			expected: `{"op": "replace", "path": "/foo~1bar~0", "value": "new value"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := patchway.Replace(tt.value, tt.path...)

			assert.JSONEq(t, tt.expected, r.String())
		})
	}
}
