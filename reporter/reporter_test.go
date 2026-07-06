package reporter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/d2202/envcheck/checker"
)

func TestReporter(t *testing.T) {
	var tests = []struct {
		name string
		in   checker.CheckResult
		want string
	}{
		{"no issues",
			checker.CheckResult{},
			"OK: no issues found",
		},
		{"missing",
			checker.CheckResult{
				Issues: []checker.Issue{{Key: "asd", Kind: checker.Missing}},
			},
			"missing keys.", // Result: %d missing keys.
		},
		{"extra",
			checker.CheckResult{
				Issues: []checker.Issue{{Key: "asd", Kind: checker.Extra}},
			},
			"extra keys.", // Result: %d extra keys.
		},
		{"empty",
			checker.CheckResult{
				Issues: []checker.Issue{{Key: "asd", Kind: checker.Empty}},
			},
			"empty keys.", // Result: %d empty keys.
		},
		{"missing > extra",
			checker.CheckResult{
				Issues: []checker.Issue{
					{Key: "asd", Kind: checker.Empty},
					{Key: "zxc", Kind: checker.Missing},
				},
			},
			"missing keys.", // Result: %d missing keys., missing - приоритетнее
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			Report(ReportInput{
				Result: tt.in,
				//len и пути не важны
				ExamplePath:    "path/to/.example.env",
				EnvPath:        "path/to/.env",
				ExampleKeysLen: 2,
				EnvKeysLen:     2,
			}, &buf)
			got := buf.String()
			if !strings.Contains(got, tt.want) {
				t.Errorf("test reporter: got %q, want: %q", got, tt.want)
			}
		})
	}
}
