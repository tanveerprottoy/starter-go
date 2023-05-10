package stringspkg

import "strings"

func Split(v, sep string) []string {
	return strings.Split(v, " ")
}

func Trim(v, cutset string) string {
	return strings.Trim(v, cutset)
}

func TrimSpace(v string) string {
	return strings.TrimSpace(v)
}

func TrimSuffix(v, suffix string) string {
	return strings.TrimSuffix(v, suffix)
}

// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func ToLower(v string) string {
	return strings.ToLower(v)
}

func ToUpper(v string) string {
	return strings.ToUpper(v)
}

func Substring(value string, start int, end int) string {
	return value[start:end]
}

func HasPrefix(value, prefix string) bool {
	return strings.HasPrefix(value, prefix)
}