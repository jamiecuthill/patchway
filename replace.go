package patchway

import (
	"encoding/json"
	"strings"
)

// Replace creates a new replace operation
//
// Example:
//
//	patchway.Replace("value", []string{"path", "to", "field"})"
func Replace(value any, path ...string) ReplaceOperation {
	for i := range path {
		path[i] = escapeJSONPointer(path[i])
	}

	return ReplaceOperation{
		Op:    "replace",
		Path:  "/" + strings.Join(path, "/"),
		Value: value,
	}
}

type ReplaceOperation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value any    `json:"value"`
}

func (o ReplaceOperation) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func escapeJSONPointer(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "~", "~0"), "/", "~1")
}
