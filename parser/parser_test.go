package parser

import (
	"bufio"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	var tests = []struct {
		name               string
		in                 string
		wantKey, wantValue string
		ok                 bool
	}{
		{"basic success", "key=value", "key", "value", true},
		{"value with embedded equals", "KEY=postgres://user:pass@host?sslmode=require", "KEY", "postgres://user:pass@host?sslmode=require", true},
		{"blank line + comment", "# comment", "", "", false},
		{"blank line + comment w/ indent", " # comment with indent", "", "", false},
		{"blank line", "", "", "", false},
		{"blank line w/ space", " ", "", "", false},
		{"inline comment", "somekey=val # this is a comment", "somekey", "val", true},
		{"export prefix", "export EXPORTKEY=exportvalue", "EXPORTKEY", "exportvalue", true},
		{"quotted value w/ spaces", "quotted=\"value with spaces\"", "quotted", "value with spaces", true},
		{"quotted value #2 w/ spaces", "quottedone='value with spaces'", "quottedone", "value with spaces", true},
		{"symbol in value", "testkey=value#value", "testkey", "value#value", true},
		{"one quote", "KEY=\"value", "KEY", "\"value", true},
		{"only key blank value", "KEY=", "KEY", "", true},
		{"random string", "asdwdasd", "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, v, ok := parseLine(tt.in)
			if k != tt.wantKey || v != tt.wantValue || ok != tt.ok {
				t.Errorf("parseLine(%q) = (%q, %q, %t), want (%q, %q, %t)",
					tt.in, k, v, ok, tt.wantKey, tt.wantValue, tt.ok)
			}
		})
	}
}

func TestParse(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, ".env")
		content := "KEY1=value1\nKEY2=value2\n"
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}

		result, err := Parse(path)
		if err != nil {
			t.Errorf("want err == nil, got: %v", err)
		}
		want := EnvMap{"KEY1": "value1", "KEY2": "value2"}
		if !maps.Equal(want, result) {
			t.Errorf("want: %q, got: %q", want, result)
		}
	})
	t.Run("file not found", func(t *testing.T) {
		result, err := Parse("path/to/file/.env")
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

		_, err := Parse(path)
		if err == nil {
			t.Error("want err != nil")
		}
	})
}
