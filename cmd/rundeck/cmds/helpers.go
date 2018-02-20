package cmds

import (
	"fmt"
	"strings"
)

// ParseSliceKeyValue parses a cobra StringSlice into a map[string]string split on an = sign
func ParseSliceKeyValue(s []string) (map[string]string, error) {
	res := map[string]string{}
	for _, o := range s {
		entry := strings.Split(o, "=")
		if len(entry) != 2 {
			return res, fmt.Errorf("error parsing entry: %s", o)
		}
		if entry[0] == "" {
			return res, fmt.Errorf("unable to parse key for value: %s", entry[1])
		}
		if entry[1] == "" {
			return res, fmt.Errorf("unable to parse value for key: %s", entry[0])
		}
		res[entry[0]] = entry[1]
	}
	return res, nil
}
