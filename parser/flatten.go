package parser

import "fmt"

func flatten(prefix string, raw map[string]any, result EnvMap) {
	for k, v := range raw {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch val := v.(type) {
		case map[string]any:
			flatten(key, val, result)
		default:
			result[key] = fmt.Sprintf("%v", val)
		}
	}
}
