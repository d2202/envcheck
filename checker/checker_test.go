package checker

import (
	"slices"
	"strings"
	"testing"

	"github.com/d2202/envcheck/parser"
)

func TestChecker(t *testing.T) {
	var tests = []struct {
		name      string
		inExample parser.EnvMap
		inEnv     parser.EnvMap
		want      CheckResult
	}{
		{"valid",
			parser.EnvMap{"KEY": "value"},
			parser.EnvMap{"KEY": "value"},
			CheckResult{},
		},
		{
			"extra",
			parser.EnvMap{"KEY": "value"},
			parser.EnvMap{"KEY": "value", "KEY2": "val2"},
			CheckResult{Issues: []Issue{{Key: "KEY2", Kind: Extra}}},
		},
		{
			"missing",
			parser.EnvMap{"KEY": "value", "KEY2": "val2"},
			parser.EnvMap{"KEY": "value"},
			CheckResult{Issues: []Issue{{Key: "KEY2", Kind: Missing}}},
		},
		{
			"empty",
			parser.EnvMap{"KEY": ""},
			parser.EnvMap{"KEY": ""},
			CheckResult{Issues: []Issue{{Key: "KEY", Kind: Empty}}},
		},
		{
			"empty envmaps",
			parser.EnvMap{},
			parser.EnvMap{},
			CheckResult{},
		},
	}
	sortIssues := func(issues []Issue) {
		slices.SortFunc(issues, func(a, b Issue) int {
			return strings.Compare(string(a.Key), string(b.Key))
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Check(tt.inExample, tt.inEnv)

			sortIssues(got.Issues)
			sortIssues(tt.want.Issues)
			if !slices.Equal(got.Issues, tt.want.Issues) {
				t.Errorf("test checker got: %v, want: %v", got.Issues, tt.want.Issues)
			}
		})
	}
}
