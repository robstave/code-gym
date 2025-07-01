// util.go
package util

import (
	"slices"
)

// Reverse returns the input string with characters in reverse order.
func Reverse1(s string) string {

	runes := []byte(s)
	slices.Reverse(runes)
	return string(runes)

}
