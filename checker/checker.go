package checker

import "github.com/d2202/envcheck/parser"

type IssueKind string

const (
	Missing IssueKind = "missing"
	Extra   IssueKind = "extra"
	Empty   IssueKind = "empty"
)

type Issue struct {
	Key  string
	Kind IssueKind
}

type CheckResult struct {
	Issues []Issue
}

func Check(example, env parser.EnvMap) CheckResult {
	var result CheckResult
	for k := range example {
		v, ok := env[k]
		if !ok {
			result.Issues = append(result.Issues, Issue{Key: k, Kind: Missing})
			continue
		}
		if v == "" {
			result.Issues = append(result.Issues, Issue{Key: k, Kind: Empty})
		}
	}
	for k := range env {
		if _, ok := example[k]; !ok {
			result.Issues = append(result.Issues, Issue{Key: k, Kind: Extra})
		}
	}
	return result
}
