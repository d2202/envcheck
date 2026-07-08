package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/d2202/envcheck/checker"
	"github.com/d2202/envcheck/parser"
	"github.com/d2202/envcheck/reporter"
)

func exitCode(result checker.CheckResult, strict bool) int {
	code := 0
	for _, issue := range result.Issues {
		if issue.Kind == checker.Missing {
			return 1
		}
		if !strict && (issue.Kind == checker.Extra || issue.Kind == checker.Empty) {
			code = 2
		}
	}
	return code
}

func mustParse(path, format string) parser.EnvMap {
	keys, err := parser.Parse(path, parser.Format(format))
	if err != nil {
		fmt.Fprintf(os.Stderr, "envcheck: %v\n", err)
		os.Exit(1)
	}
	return keys
}

func main() {
	envPath := flag.String("actual", ".env", "path to actual file")
	examplePath := flag.String("expected", ".env.example", "path to expected (.example) file")
	quiet := flag.Bool("quiet", false, "only exit code, no output")
	strict := flag.Bool("strict", false, "only missing keys count as errors")
	writeJSON := flag.Bool("json", false, "format result in JSON")
	format := flag.String("format", "env", "compare files format")

	flag.Parse()

	exampleKeys := mustParse(*examplePath, *format)
	envKeys := mustParse(*envPath, *format)

	checkRes := checker.Check(exampleKeys, envKeys)
	input := reporter.ReportInput{
		Result:         checkRes,
		ExamplePath:    *examplePath,
		EnvPath:        *envPath,
		ExampleKeysLen: len(exampleKeys),
		EnvKeysLen:     len(envKeys),
	}
	if !*quiet {
		if *writeJSON {
			err := reporter.ReportJSON(input, os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "envcheck: failed to write JSON output\n")
				os.Exit(1)
			}
		} else {
			reporter.Report(input, os.Stdout)
		}
	}
	os.Exit(exitCode(checkRes, *strict))
}
