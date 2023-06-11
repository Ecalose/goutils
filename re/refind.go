package re

import (
	"regexp"
)

func ReFindList(str, rule string) []string {
	compileRegex := regexp.MustCompile(rule).FindStringSubmatch(str)
	return compileRegex[1:]
}
func ReFindString(str, rule string) string {
	compileRegex := regexp.MustCompile(rule).FindString(str)
	return compileRegex
}
