package reporter

import (
	"fmt"
	"io"

	"github.com/d2202/envcheck/checker"
)

type ReportInput struct {
	Result         checker.CheckResult
	ExamplePath    string
	EnvPath        string
	ExampleKeysLen int
	EnvKeysLen     int
}

func categorize(issues []checker.Issue) (missing, extra, empty []string) {
	for _, issue := range issues {
		switch issue.Kind {
		case checker.Missing:
			missing = append(missing, issue.Key)
		case checker.Extra:
			extra = append(extra, issue.Key)
		case checker.Empty:
			empty = append(empty, issue.Key)
		}
	}
	return
}

func Report(in ReportInput, w io.Writer) {
	if len(in.Result.Issues) == 0 {
		fmt.Fprintln(w, "✅\tOK: no issues found.")
		return
	}

	fmt.Fprintf(w,
		"Comparing:\n\texample\t: %s (%d keys)\n\tenv\t: %s (%d keys)\n",
		in.ExamplePath,
		in.ExampleKeysLen,
		in.EnvPath,
		in.EnvKeysLen,
	)

	missing, extra, empty := categorize(in.Result.Issues)
	if len(extra) != 0 {
		fmt.Fprintf(w, "⚠️\tEXTRA (%d)\t— keys in env, not in example\n", len(extra))
		for _, v := range extra {
			fmt.Fprintln(w, v)
		}
	}

	if len(empty) != 0 {
		fmt.Fprintf(w, "⚠️\tEMPTY (%d)\t— key exists but value is blank\n", len(empty))
		for _, v := range empty {
			fmt.Fprintln(w, v)
		}
	}

	if len(missing) != 0 {
		fmt.Fprintf(w, "❌\tMISSING (%d)\t— keys in example, not in env\n", len(missing))
		for _, v := range missing {
			fmt.Fprintln(w, v)
		}
	}

	switch {
	case len(missing) != 0:
		fmt.Fprintf(w, "Result: %d missing keys.\n", len(missing))
	case len(empty) != 0:
		fmt.Fprintf(w, "Result: %d empty keys.\n", len(empty))
	case len(extra) != 0:
		fmt.Fprintf(w, "Result: %d extra keys.\n", len(extra))
	}
}
