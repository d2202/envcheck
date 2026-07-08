package parser

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestParseToml(t *testing.T) {
	var tests = []struct {
		name    string
		content string
		want    EnvMap
	}{
		{
			"basic parsing",
			"[database]\nhost = \"localhost\"\nport = 5432",
			EnvMap{"database.host": "localhost", "database.port": "5432"},
		},
		{
			"blank file",
			"",
			EnvMap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "test.toml")
			err := os.WriteFile(path, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("parse toml: %v", err)
			}
			got, err := parseToml(path)
			if err != nil {
				t.Errorf("parse toml: want err == nil, got: %v", err)
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("parse toml: want %v, got: %v", got, tt.want)
			}
		})
	}
}
