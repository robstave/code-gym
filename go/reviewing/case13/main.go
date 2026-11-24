// Case 13: Remove a prefix. Difficulty: Easy.
package main

import (
	"fmt"
	"strings"
)

func removeLogPrefix(label string) string {
	return strings.TrimLeft(label, "log-")
}

func main() {
	fmt.Println("Label:", removeLogPrefix("log-golf"))
}
