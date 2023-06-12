package re

import (
	"regexp"
)

// 类似python的findall
func FindList(str, rule string) []string {
	compileRegex := regexp.MustCompile(rule).FindStringSubmatch(str)
	return compileRegex[1:]
}
func FindString(str, rule string) string {
	compileRegex := regexp.MustCompile(rule).FindString(str)
	return compileRegex
}
