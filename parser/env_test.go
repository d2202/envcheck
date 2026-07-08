package parser

import (
	"maps"
	"os"
	"path/filepath"
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

func TestParseEnv(t *testing.T) {
	var tests = []struct {
		name    string
		content string
		want    EnvMap
	}{
		{
			"basic parsing",
			"KEY=val\nOTHERKEY=otherval",
			EnvMap{"KEY": "val", "OTHERKEY": "otherval"},
		},
		{
			"blank file",
			"",
			EnvMap{},
		},
		{
			"parsing bad key",
			"KEY=val\nOTHERKEY=otherval\n//comment\nBLAHBLAH",
			EnvMap{"KEY": "val", "OTHERKEY": "otherval"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, ".env")
			err := os.WriteFile(path, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("parse env: %v", err)
			}
			got, err := parseEnv(path)
			if err != nil {
				t.Errorf("parse env: want err == nil, got: %v", err)
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("parse env: want %v, got: %v", got, tt.want)
			}
		})
	}
}
