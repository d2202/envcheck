package parser

import (
	"maps"
	"testing"
)

func TestFlatten(t *testing.T) {
	raw := map[string]any{
		"database": map[string]any{
			"host": "localhost",
			"credentials": map[string]any{
				"user": "admin",
			},
		},
		"debug": true,
	}
	want := EnvMap{
		"database.host":             "localhost",
		"database.credentials.user": "admin",
		"debug":                     "true",
	}
	result := make(EnvMap)
	flatten("", raw, result)
	if !maps.Equal(result, want) {
		t.Errorf("flatten: want %v, got %v", want, result)
	}
}
