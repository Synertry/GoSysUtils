package Str

import "strings"

// Concat is an efficient string builder function
func Concat(str ...string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}
