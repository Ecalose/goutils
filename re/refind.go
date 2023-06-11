package re

import (
	"regexp"
)

// 类似findall
func ReFindList(str, rule string) []string {
	compileRegex := regexp.MustCompile(rule).FindStringSubmatch(str)
	return compileRegex[1:]
}
func ReFindString(str, rule string) string {
	compileRegex := regexp.MustCompile(rule).FindString(str)
	return compileRegex
}
