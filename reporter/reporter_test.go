package reporter

import (
	"bytes"
	"encoding/json"
	"slices"
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

func TestReportJSON(t *testing.T) {
	var tests = []struct {
		name string
		in   checker.CheckResult
		want JSONReport
	}{
		{"no issues",
			checker.CheckResult{},
			JSONReport{OK: true},
		},
		{"missing",
			checker.CheckResult{
				Issues: []checker.Issue{{Key: "asd", Kind: checker.Missing}},
			},
			JSONReport{Missing: []string{"asd"}},
		},
		{"extra",
			checker.CheckResult{
				Issues: []checker.Issue{{Key: "asd", Kind: checker.Extra}},
			},
			JSONReport{Extra: []string{"asd"}},
		},
		{"empty",
			checker.CheckResult{
				Issues: []checker.Issue{{Key: "asd", Kind: checker.Empty}},
			},
			JSONReport{Empty: []string{"asd"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				buf bytes.Buffer
				got JSONReport
			)

			ReportJSON(ReportInput{
				Result: tt.in,
				//len и пути не важны
				ExamplePath:    "path/to/.example.env",
				EnvPath:        "path/to/.env",
				ExampleKeysLen: 2,
				EnvKeysLen:     2,
			}, &buf)

			if err := json.NewDecoder(&buf).Decode(&got); err != nil {
				t.Fatalf("failed to decode JSON: %v", err)
			}

			if got.OK != tt.want.OK {
				t.Errorf("test report json: got: %t, want: %t", got.OK, tt.want.OK)
			}

			compareResultFields(t, got.Missing, tt.want.Missing, "missing")
			compareResultFields(t, got.Empty, tt.want.Empty, "empty")
			compareResultFields(t, got.Extra, tt.want.Extra, "extra")
		})
	}
}

func compareResultFields(t *testing.T, got, want []string, field string) {
	t.Helper()
	slices.Sort(got)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Errorf("%s: got %v, want %v", field, got, want)
	}
}
