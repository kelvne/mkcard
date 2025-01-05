package helpers

import (
	"regexp"
	"strings"
)

// SnakeCase converts a string to snake_case, if a space is found
// it behaves differently than if it is not found. This is the
// intended behavior. See the examples for more information.
func SnakeCase(value string) string {
	str := strings.TrimSpace(value)
	str = strings.ReplaceAll(str, " ", "_")
	if !strings.Contains(value, " ") {
		re := regexp.MustCompile("([a-z0-9])([A-Z])")
		str = re.ReplaceAllString(str, "${1}_${2}")
		re = regexp.MustCompile("([A-Z])([A-Z][a-z])")
		str = re.ReplaceAllString(str, "${1}_${2}")
	}
	return strings.ToLower(str)
}
