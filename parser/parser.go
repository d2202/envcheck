package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type EnvMap map[string]string

func Parse(path string) (result EnvMap, err error) {
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
		line := scanner.Text()
		kvPair := strings.SplitN(line, "=", 2)
		if len(kvPair) > 1 {
			// строка без "=" — пропускаем (нет валидного KV)
			result[kvPair[0]] = kvPair[1]
		}
	}
	if scErr := scanner.Err(); scErr != nil {
		return nil, fmt.Errorf("parse: %w", scErr)
	}
	return result, nil
}
