package dotted

import (
	"fmt"
	"strings"
)

// ToMap ...
func ToMap(s string) (map[string]interface{}, error) {
	kv := strings.Split(s, "=")
	if len(kv) != 2 {
		return nil, fmt.Errorf("invalid value for set variable: it must be path.to.key=value ")
	}

	keys := strings.Split(kv[0], ".")
	m := map[string]interface{}{}
	var r map[string]interface{}
	r = m
	for i, k := range keys {
		if i != len(keys)-1 {
			r[k] = map[string]interface{}{}
			r = r[k].(map[string]interface{})
		} else {
			r[k] = kv[1]
		}
	}
	return m, nil
}
