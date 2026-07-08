package parser

import (
	"fmt"
)

type EnvMap map[string]string

type Format string

const (
	Env  Format = "env"
	Toml Format = "toml"
	Yaml Format = "yaml"
)

func Parse(path string, format Format) (result EnvMap, err error) {
	switch format {
	case Env:
		return parseEnv(path)
	case Toml:
		return parseToml(path)
	case Yaml:
		return parseYaml(path)
	default:
		return nil, fmt.Errorf("parse: unsupported format: %s", format)
	}
}
