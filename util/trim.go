package util

import (
	"strings"
)

// Removes white space from the beginning and end of a string,
// including spaces, tabs and new lines.
func Trim(str string) string {
	return strings.Trim(str, " \t\r\n")
}
