package parser

import (
	"bufio"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParse_FileExtensions(t *testing.T) {
	var tests = []struct {
		name    string
		ext     string
		format  Format
		content string
		want    EnvMap
	}{
		{
			"env",
			".env",
			Env,
			"KEY1=value1\nKEY2=value2\n",
			EnvMap{"KEY1": "value1", "KEY2": "value2"},
		},
		{
			"toml",
			".toml",
			Toml,
			"[database]\nhost = \"localhost\"\nport = 5432",
			EnvMap{"database.host": "localhost", "database.port": "5432"},
		},
		{
			"yaml",
			".yaml",
			Yaml,
			"database:\n  host: localhost\n  port: 5432",
			EnvMap{"database.host": "localhost", "database.port": "5432"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, tt.ext)
			if err := os.WriteFile(path, []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}

			result, err := Parse(path, tt.format)
			if err != nil {
				t.Errorf("want err == nil, got: %v", err)
			}
			if !maps.Equal(tt.want, result) {
				t.Errorf("want: %q, got: %q", tt.want, result)
			}
		})
	}
}

func TestParse(t *testing.T) {
	t.Run("file not found", func(t *testing.T) {
		result, err := Parse("path/to/file/.env", Env)
		if err == nil {
			t.Error("want err != nil, got nil")
		}
		if result != nil {
			t.Errorf("want result == nil, got %v", result)
		}
	})

	t.Run("line too long", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, ".env")
		longLine := "KEY=" + strings.Repeat("x", bufio.MaxScanTokenSize+1)
		os.WriteFile(path, []byte(longLine), 0644)

		_, err := Parse(path, Env)
		if err == nil {
			t.Error("want err != nil")
		}
	})

	t.Run("unsupported file format", func(t *testing.T) {
		result, err := Parse("path/to/file/.conf", Format("conf"))
		if err == nil {
			t.Error("want err != nil, got nil")
		}
		if result != nil {
			t.Errorf("want result == nil, got %v", result)
		}
	})
}
