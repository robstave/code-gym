// Case 16: Missing account. Difficulty: Medium.
package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("account not found")

func lookupAccount(id string) error {
	return fmt.Errorf("lookup %q: %w", id, errNotFound)
}

func main() {
	err := lookupAccount("123")
	if err == errNotFound {
		fmt.Println("Account not found")
	} else if err != nil {
		fmt.Println("Lookup failed")
	}
}
