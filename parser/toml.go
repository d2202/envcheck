package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func parseToml(path string) (result EnvMap, err error) {
	result = make(EnvMap)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("parse toml %s: %w", path, err)
	}
	defer func() {
		if clErr := file.Close(); clErr != nil {
			err = errors.Join(err, clErr)
		}
	}()
	var raw map[string]any
	if _, err := toml.NewDecoder(file).Decode(&raw); err != nil {
		return nil, fmt.Errorf("parse toml %s: %w", path, err)
	}
	flatten("", raw, result)
	return result, nil
}
