package xstring

import "strings"

func ExcludeWhiteSpaces(v string) string {
	return strings.Join(strings.Fields(v), "")
}
