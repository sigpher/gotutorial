package funcpack

import "strings"

func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + "." + suffix
	}
}
