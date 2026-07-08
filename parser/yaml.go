package parser

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func parseYaml(path string) (result EnvMap, err error) {
	result = make(EnvMap)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("parse yaml %s: %w", path, err)
	}
	defer func() {
		if clErr := file.Close(); clErr != nil {
			err = errors.Join(err, clErr)
		}
	}()
	var raw map[string]any
	if err := yaml.NewDecoder(file).Decode(&raw); err != nil {
		if errors.Is(err, io.EOF) {
			return result, nil
		}
		return nil, fmt.Errorf("parse yaml %s: %w", path, err)
	}
	flatten("", raw, result)
	return result, nil
}
