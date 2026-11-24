// Case 8: First character. Difficulty: Medium.
package main

import (
	"fmt"
	"unicode/utf8"
)

func firstCharacter(text string) string {
	return text[:1]
}

func main() {
	first := firstCharacter("éclair")
	fmt.Println("Valid UTF-8:", utf8.ValidString(first))
}
