// util.go
package util

import (
	"slices"
)

// Reverse returns the input string with characters in reverse order.
func Reverse1(s string) string {

	chars := []byte(s)
	slices.Reverse(chars)
	return string(chars)

}

func Reverse2(s string) string {

	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)

}
