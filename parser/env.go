package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const exportPrefix string = "export "

func parseEnv(path string) (result EnvMap, err error) {
	result = make(EnvMap)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	defer func() {
		if clErr := file.Close(); clErr != nil {
			err = errors.Join(err, clErr)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if key, value, ok := parseLine(scanner.Text()); ok {
			result[key] = value
		}
	}
	if scErr := scanner.Err(); scErr != nil {
		return nil, fmt.Errorf("parse env: %w", scErr)
	}
	return result, nil
}

// parseLine разбирает одну строку .env-файла.
// ok=false означает что строку нужно пропустить (пустая, комментарий, без "=").
func parseLine(line string) (key, value string, ok bool) {
	line = strings.TrimSpace(line)

	if !shouldProcessLine(line) {
		return "", "", false
	}

	kvPair := strings.SplitN(line, "=", 2)
	if len(kvPair) < 2 {
		return "", "", false
	}

	key, value = prepareKeyValue(kvPair[0], kvPair[1])
	return key, value, true
}

func shouldProcessLine(line string) bool {
	return line != "" && !strings.HasPrefix(line, "#")
}

func isQuotedString(s string) bool {
	if len(s) < 2 {
		return false
	}
	quote := s[0]
	return (quote == '"' || quote == '\'') && s[len(s)-1] == quote
}

func prepareKeyValue(key, value string) (string, string) {
	key, value = strings.TrimSpace(key), strings.TrimSpace(value)
	resultKey, _ := strings.CutPrefix(key, exportPrefix)

	if isQuotedString(value) {
		return resultKey, value[1 : len(value)-1]
	}

	resultValue, _, _ := strings.Cut(value, " ")
	return resultKey, resultValue
}
