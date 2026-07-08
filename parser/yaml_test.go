package parser

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestParseYaml(t *testing.T) {
	var tests = []struct {
		name    string
		content string
		want    EnvMap
	}{
		{
			"basic parsing",
			"database:\n  host: localhost\n  port: 5432\n  credentials:\n    user: admin\n    password: secret",
			EnvMap{
				"database.host":                 "localhost",
				"database.port":                 "5432",
				"database.credentials.user":     "admin",
				"database.credentials.password": "secret",
			},
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
			path := filepath.Join(dir, "test.yaml")
			err := os.WriteFile(path, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("parse yaml: %v", err)
			}
			got, err := parseYaml(path)
			if err != nil {
				t.Errorf("parse yaml: want err == nil, got: %v", err)
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("parse yaml: want %v, got: %v", got, tt.want)
			}
		})
	}
}
