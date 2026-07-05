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

func mustParse(path string) parser.EnvMap {
	keys, err := parser.Parse(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "envcheck: %v\n", err)
		os.Exit(1)
	}
	return keys
}

func main() {
	envPath := flag.String("env", ".env", "path to .env file")
	examplePath := flag.String("example", ".env.example", "path to .env.example file")
	quiet := flag.Bool("quiet", false, "only exit code, no output")
	strict := flag.Bool("strict", false, "only missing keys count as errors")

	flag.Parse()

	exampleKeys := mustParse(*examplePath)
	envKeys := mustParse(*envPath)

	checkRes := checker.Check(exampleKeys, envKeys)
	if !*quiet {
		reporter.Report(reporter.ReportInput{
			Result:         checkRes,
			ExamplePath:    *examplePath,
			EnvPath:        *envPath,
			ExampleKeysLen: len(exampleKeys),
			EnvKeysLen:     len(envKeys),
		}, os.Stdout)
	}
	os.Exit(exitCode(checkRes, *strict))
}
