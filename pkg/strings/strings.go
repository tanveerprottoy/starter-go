package strings

import "strings"

func Split(v, sep string) []string {
	return strings.Split(v, " ")
}

func Trim(v, cutset string) string {
	return strings.Trim(v, cutset)
}

func ToLower(v string) string {
	return strings.ToLower(v)
}

func ToUpper(v string) string {
	return strings.ToUpper(v)
}
